import json
import logging
import uuid
from datetime import datetime
from typing import Dict, Any, Callable, Optional
import pika
from pika.adapters.blocking_connection import BlockingChannel
from pika.exceptions import AMQPConnectionError
import time

logger = logging.getLogger(__name__)

class RabbitMQClient:
    """RabbitMQ client for publishing and consuming payment events"""
    
    def __init__(self, url: str):
        self.url = url
        self.connection = None
        self.channel = None
    
    def connect(self) -> bool:
        """Establish connection to RabbitMQ with retry logic"""
        max_retries = 5
        for attempt in range(max_retries):
            try:
                # Parse connection parameters
                parameters = pika.URLParameters(self.url)
                self.connection = pika.BlockingConnection(parameters)
                self.channel = self.connection.channel()
                
                # Setup infrastructure
                self._setup_infrastructure()
                
                logger.info("Successfully connected to RabbitMQ")
                return True
                
            except AMQPConnectionError as e:
                logger.warning(f"Failed to connect to RabbitMQ (attempt {attempt + 1}): {e}")
                if attempt < max_retries - 1:
                    time.sleep(attempt + 1)  # Exponential backoff
                
        logger.error("Failed to connect to RabbitMQ after all retries")
        return False
    
    def _setup_infrastructure(self) -> None:
        """Setup exchanges and queues"""
        # Declare main exchange
        self.channel.exchange_declare(
            exchange='marketplace.events',
            exchange_type='topic',
            durable=True
        )
        
        # Declare payment events queue
        self.channel.queue_declare(
            queue='payment.events',
            durable=True
        )
        
        # Bind payment events queue to exchange
        self.channel.queue_bind(
            exchange='marketplace.events',
            queue='payment.events',
            routing_key='payment.events.*'
        )
        
        # Declare notification events queue (for payment notifications)
        self.channel.queue_declare(
            queue='notification.events',
            durable=True
        )
        
        # Bind notification queue to payment events
        self.channel.queue_bind(
            exchange='marketplace.events',
            queue='notification.events',
            routing_key='payment.events.*'
        )
    
    def _create_event(self, event_type: str, data: Dict[str, Any], 
                     correlation_id: str, user_id: str) -> Dict[str, Any]:
        """Create a standardized event structure"""
        return {
            "eventId": str(uuid.uuid4()),
            "eventType": event_type,
            "timestamp": datetime.utcnow().isoformat() + "Z",
            "version": "1.0.0",
            "data": data,
            "metadata": {
                "source": "payment-service",
                "correlationId": correlation_id,
                "userId": user_id
            }
        }
    
    def publish_payment_initiated(self, payment_data: Dict[str, Any], 
                                correlation_id: str, user_id: str) -> bool:
        """Publish payment initiated event"""
        event = self._create_event(
            event_type="payment.initiated",
            data={
                "paymentId": payment_data["payment_id"],
                "orderId": payment_data.get("order_id"),
                "userId": payment_data["user_id"],
                "amount": payment_data["amount"],
                "currency": payment_data["currency"],
                "provider": payment_data["provider"],
                "status": payment_data["status"],
                "createdAt": payment_data.get("created_at", datetime.utcnow().isoformat() + "Z")
            },
            correlation_id=correlation_id,
            user_id=user_id
        )
        
        return self._publish_event("payment.events.initiated", event)
    
    def publish_payment_completed(self, payment_data: Dict[str, Any], 
                                transaction_id: str, correlation_id: str, user_id: str) -> bool:
        """Publish payment completed event"""
        event = self._create_event(
            event_type="payment.completed",
            data={
                "paymentId": payment_data["payment_id"],
                "orderId": payment_data.get("order_id"),
                "userId": payment_data["user_id"],
                "amount": payment_data["amount"],
                "currency": payment_data["currency"],
                "provider": payment_data["provider"],
                "transactionId": transaction_id,
                "completedAt": datetime.utcnow().isoformat() + "Z"
            },
            correlation_id=correlation_id,
            user_id=user_id
        )
        
        return self._publish_event("payment.events.completed", event)
    
    def publish_payment_failed(self, payment_data: Dict[str, Any], 
                             error_code: str, error_message: str,
                             correlation_id: str, user_id: str) -> bool:
        """Publish payment failed event"""
        event = self._create_event(
            event_type="payment.failed",
            data={
                "paymentId": payment_data["payment_id"],
                "orderId": payment_data.get("order_id"),
                "userId": payment_data["user_id"],
                "amount": payment_data["amount"],
                "currency": payment_data["currency"],
                "provider": payment_data["provider"],
                "errorCode": error_code,
                "errorMessage": error_message,
                "failedAt": datetime.utcnow().isoformat() + "Z"
            },
            correlation_id=correlation_id,
            user_id=user_id
        )
        
        return self._publish_event("payment.events.failed", event)
    
    def _publish_event(self, routing_key: str, event: Dict[str, Any]) -> bool:
        """Publish an event to RabbitMQ"""
        try:
            if not self.channel:
                logger.error("No active channel to publish event")
                return False
                
            self.channel.basic_publish(
                exchange='marketplace.events',
                routing_key=routing_key,
                body=json.dumps(event),
                properties=pika.BasicProperties(
                    content_type='application/json',
                    delivery_mode=2,  # Make message persistent
                    message_id=event["eventId"],
                    timestamp=int(datetime.now().timestamp()),
                    headers={
                        "eventType": event["eventType"],
                        "source": event["metadata"]["source"],
                        "correlationId": event["metadata"]["correlationId"]
                    }
                )
            )
            
            logger.info(f"Published event: {event['eventType']} with routing key: {routing_key}")
            return True
            
        except Exception as e:
            logger.error(f"Failed to publish event: {e}")
            return False
    
    def consume_events(self, queue_name: str, callback: Callable[[Dict[str, Any]], bool]) -> None:
        """Start consuming events from a queue"""
        def wrapper(ch: BlockingChannel, method, properties, body):
            try:
                event = json.loads(body.decode('utf-8'))
                
                if callback(event):
                    ch.basic_ack(delivery_tag=method.delivery_tag)
                else:
                    # Requeue for retry
                    ch.basic_nack(delivery_tag=method.delivery_tag, requeue=True)
                    
            except json.JSONDecodeError as e:
                logger.error(f"Failed to decode event: {e}")
                ch.basic_nack(delivery_tag=method.delivery_tag, requeue=False)
            except Exception as e:
                logger.error(f"Failed to process event: {e}")
                ch.basic_nack(delivery_tag=method.delivery_tag, requeue=True)
        
        if not self.channel:
            logger.error("No active channel to consume events")
            return
            
        self.channel.basic_qos(prefetch_count=1)  # Process one message at a time
        self.channel.basic_consume(
            queue=queue_name,
            on_message_callback=wrapper
        )
        
        logger.info(f"Started consuming events from queue: {queue_name}")
        self.channel.start_consuming()
    
    def close(self) -> None:
        """Close the RabbitMQ connection"""
        if self.channel and not self.channel.is_closed:
            self.channel.close()
        if self.connection and not self.connection.is_closed:
            self.connection.close()
    
    def is_connected(self) -> bool:
        """Check if the client is connected"""
        return (self.connection is not None and 
                not self.connection.is_closed and 
                self.channel is not None and 
                not self.channel.is_closed)


# Event handler classes for type safety
class PaymentEventHandler:
    """Base class for handling payment events"""
    
    def handle_user_created(self, event: Dict[str, Any]) -> bool:
        """Handle user created event - could create payment profile"""
        logger.info(f"Received user created event: {event['data']['userId']}")
        # Implement payment profile creation logic here
        return True
    
    def handle_order_created(self, event: Dict[str, Any]) -> bool:
        """Handle order created event - prepare for payment"""
        logger.info(f"Received order created event: {event['data']['orderId']}")
        # Implement payment preparation logic here
        return True
    
    def handle_event(self, event: Dict[str, Any]) -> bool:
        """Route events to appropriate handlers"""
        event_type = event.get("eventType")
        
        if event_type == "user.created":
            return self.handle_user_created(event)
        elif event_type == "order.created":
            return self.handle_order_created(event)
        else:
            logger.warning(f"Unhandled event type: {event_type}")
            return True  # Acknowledge unknown events to prevent reprocessing 