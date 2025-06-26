package messaging

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

// RabbitMQClient handles RabbitMQ operations
type RabbitMQClient struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	url        string
}

// Event represents a marketplace event
type Event struct {
	EventID   string                 `json:"eventId"`
	EventType string                 `json:"eventType"`
	Timestamp string                 `json:"timestamp"`
	Version   string                 `json:"version"`
	Data      map[string]interface{} `json:"data"`
	Metadata  EventMetadata          `json:"metadata"`
}

// EventMetadata contains event metadata
type EventMetadata struct {
	Source        string `json:"source"`
	CorrelationID string `json:"correlationId"`
	UserID        string `json:"userId"`
}

// UserCreatedEvent represents a user created event
type UserCreatedEvent struct {
	UserID    string `json:"userId"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	IsActive  bool   `json:"isActive"`
	CreatedAt string `json:"createdAt"`
}

// UserUpdatedEvent represents a user updated event
type UserUpdatedEvent struct {
	UserID    string                 `json:"userId"`
	Email     string                 `json:"email"`
	FirstName string                 `json:"firstName"`
	LastName  string                 `json:"lastName"`
	IsActive  bool                   `json:"isActive"`
	UpdatedAt string                 `json:"updatedAt"`
	Changes   map[string]interface{} `json:"changes"`
}

// NewRabbitMQClient creates a new RabbitMQ client
func NewRabbitMQClient(url string) *RabbitMQClient {
	return &RabbitMQClient{
		url: url,
	}
}

// Connect establishes connection to RabbitMQ
func (r *RabbitMQClient) Connect() error {
	var err error

	// Retry connection with exponential backoff
	for attempts := 0; attempts < 5; attempts++ {
		r.connection, err = amqp.Dial(r.url)
		if err == nil {
			break
		}

		log.Printf("Failed to connect to RabbitMQ (attempt %d): %v", attempts+1, err)
		time.Sleep(time.Duration(attempts+1) * time.Second)
	}

	if err != nil {
		return fmt.Errorf("failed to connect to RabbitMQ after retries: %w", err)
	}

	r.channel, err = r.connection.Channel()
	if err != nil {
		return fmt.Errorf("failed to open channel: %w", err)
	}

	// Setup infrastructure
	if err := r.setupInfrastructure(); err != nil {
		return fmt.Errorf("failed to setup infrastructure: %w", err)
	}

	log.Println("Successfully connected to RabbitMQ")
	return nil
}

// setupInfrastructure creates exchanges and queues
func (r *RabbitMQClient) setupInfrastructure() error {
	// Declare main exchange
	err := r.channel.ExchangeDeclare(
		"marketplace.events", // name
		"topic",              // type
		true,                 // durable
		false,                // auto-deleted
		false,                // internal
		false,                // no-wait
		nil,                  // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare exchange: %w", err)
	}

	// Declare user events queue
	_, err = r.channel.QueueDeclare(
		"user.events", // name
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare user events queue: %w", err)
	}

	// Bind user events queue to exchange
	err = r.channel.QueueBind(
		"user.events",        // queue name
		"user.events.*",      // routing key pattern
		"marketplace.events", // exchange
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to bind user events queue: %w", err)
	}

	// Declare notification events queue
	_, err = r.channel.QueueDeclare(
		"notification.events", // name
		true,                  // durable
		false,                 // delete when unused
		false,                 // exclusive
		false,                 // no-wait
		nil,                   // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare notification events queue: %w", err)
	}

	// Bind notification queue to user events
	err = r.channel.QueueBind(
		"notification.events", // queue name
		"user.events.*",       // routing key pattern
		"marketplace.events",  // exchange
		false,
		nil,
	)

	return err
}

// PublishUserCreated publishes a user created event
func (r *RabbitMQClient) PublishUserCreated(userEvent UserCreatedEvent, correlationID, userID string) error {
	event := Event{
		EventID:   uuid.New().String(),
		EventType: "user.created",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Version:   "1.0.0",
		Data: map[string]interface{}{
			"userId":    userEvent.UserID,
			"email":     userEvent.Email,
			"firstName": userEvent.FirstName,
			"lastName":  userEvent.LastName,
			"isActive":  userEvent.IsActive,
			"createdAt": userEvent.CreatedAt,
		},
		Metadata: EventMetadata{
			Source:        "user-service",
			CorrelationID: correlationID,
			UserID:        userID,
		},
	}

	return r.publishEvent("user.events.created", event)
}

// PublishUserUpdated publishes a user updated event
func (r *RabbitMQClient) PublishUserUpdated(userEvent UserUpdatedEvent, correlationID, userID string) error {
	event := Event{
		EventID:   uuid.New().String(),
		EventType: "user.updated",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Version:   "1.0.0",
		Data: map[string]interface{}{
			"userId":    userEvent.UserID,
			"email":     userEvent.Email,
			"firstName": userEvent.FirstName,
			"lastName":  userEvent.LastName,
			"isActive":  userEvent.IsActive,
			"updatedAt": userEvent.UpdatedAt,
			"changes":   userEvent.Changes,
		},
		Metadata: EventMetadata{
			Source:        "user-service",
			CorrelationID: correlationID,
			UserID:        userID,
		},
	}

	return r.publishEvent("user.events.updated", event)
}

// publishEvent publishes an event to RabbitMQ
func (r *RabbitMQClient) publishEvent(routingKey string, event Event) error {
	body, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	err = r.channel.Publish(
		"marketplace.events", // exchange
		routingKey,           // routing key
		false,                // mandatory
		false,                // immediate
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         body,
			DeliveryMode: amqp.Persistent, // make message persistent
			MessageId:    event.EventID,
			Timestamp:    time.Now(),
			Headers: amqp.Table{
				"eventType":     event.EventType,
				"source":        event.Metadata.Source,
				"correlationId": event.Metadata.CorrelationID,
			},
		},
	)

	if err != nil {
		return fmt.Errorf("failed to publish event: %w", err)
	}

	log.Printf("Published event: %s with routing key: %s", event.EventType, routingKey)
	return nil
}

// ConsumeEvents starts consuming events from a queue
func (r *RabbitMQClient) ConsumeEvents(queueName string, handler func(Event) error) error {
	msgs, err := r.channel.Consume(
		queueName, // queue
		"",        // consumer
		false,     // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		return fmt.Errorf("failed to register consumer: %w", err)
	}

	go func() {
		for d := range msgs {
			var event Event
			if err := json.Unmarshal(d.Body, &event); err != nil {
				log.Printf("Failed to unmarshal event: %v", err)
				d.Nack(false, false) // don't requeue
				continue
			}

			if err := handler(event); err != nil {
				log.Printf("Failed to handle event %s: %v", event.EventID, err)
				d.Nack(false, true) // requeue for retry
			} else {
				d.Ack(false) // acknowledge successful processing
			}
		}
	}()

	log.Printf("Started consuming events from queue: %s", queueName)
	return nil
}

// Close closes the RabbitMQ connection
func (r *RabbitMQClient) Close() error {
	if r.channel != nil {
		r.channel.Close()
	}
	if r.connection != nil {
		r.connection.Close()
	}
	return nil
}

// IsConnected checks if the client is connected
func (r *RabbitMQClient) IsConnected() bool {
	return r.connection != nil && !r.connection.IsClosed()
}
