package application

import (
	"context"
	"errors"
	"time"
	"user-service/domain"
	"user-service/infrastructure/messaging"

	"github.com/google/uuid"
)

// UserRepository defines the interface for user data access
type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	GetByID(ctx context.Context, id string) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, id string) error
}

// MessagePublisher defines the interface for publishing events
type MessagePublisher interface {
	PublishUserCreated(userEvent messaging.UserCreatedEvent, correlationID, userID string) error
	PublishUserUpdated(userEvent messaging.UserUpdatedEvent, correlationID, userID string) error
}

// UserService implements the application use cases for user management
type UserService struct {
	userRepo  UserRepository
	publisher MessagePublisher
}

// NewUserService creates a new instance of UserService
func NewUserService(userRepo UserRepository, publisher MessagePublisher) *UserService {
	return &UserService{
		userRepo:  userRepo,
		publisher: publisher,
	}
}

// CreateUser creates a new user in the system
func (s *UserService) CreateUser(ctx context.Context, email, username, password string) (*domain.User, error) {
	// Validate input
	if email == "" || username == "" || password == "" {
		return nil, errors.New("email, username, and password are required")
	}

	// Check if user already exists
	existingUser, _ := s.userRepo.GetByEmail(ctx, email)
	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	// Create new user
	user := domain.NewUser(email, username, password)

	// Save to repository
	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	// Publish user created event
	correlationID := uuid.New().String()
	userCreatedEvent := messaging.UserCreatedEvent{
		UserID:    user.ID.String(),
		Email:     user.Email,
		FirstName: user.Username,
		LastName:  "",
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}

	if s.publisher != nil {
		if err := s.publisher.PublishUserCreated(userCreatedEvent, correlationID, user.ID.String()); err != nil {
			// Log error but don't fail the user creation
			// In a real system, you might want to use a different error handling strategy
			// such as storing failed events for retry
		}
	}

	return user, nil
}

// GetUser retrieves a user by their ID
func (s *UserService) GetUser(ctx context.Context, id string) (*domain.User, error) {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByID retrieves a user by their ID
func (s *UserService) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	if id == "" {
		return nil, errors.New("user ID is required")
	}

	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByEmail retrieves a user by their email address
func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	if email == "" {
		return nil, errors.New("email is required")
	}

	user, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// ValidateCredentials validates user credentials for authentication
func (s *UserService) ValidateCredentials(ctx context.Context, email, password string) (bool, *domain.User, error) {
	if email == "" || password == "" {
		return false, nil, errors.New("email and password are required")
	}

	// Get user by email
	user, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return false, nil, err
	}

	// Validate password (in real implementation, you'd hash the password and compare)
	if user.ValidatePassword(password) {
		return true, user, nil
	}

	return false, nil, nil
}

// UpdateUser updates user information
func (s *UserService) UpdateUser(ctx context.Context, updatedUser *domain.User) error {
	if updatedUser == nil {
		return errors.New("user is required")
	}

	// Get current user state for tracking changes
	currentUser, err := s.userRepo.GetByID(ctx, updatedUser.ID.String())
	if err != nil {
		return err
	}

	// Track changes
	changes := make(map[string]interface{})
	if currentUser.Email != updatedUser.Email {
		changes["email"] = map[string]interface{}{
			"old": currentUser.Email,
			"new": updatedUser.Email,
		}
	}
	if currentUser.Username != updatedUser.Username {
		changes["username"] = map[string]interface{}{
			"old": currentUser.Username,
			"new": updatedUser.Username,
		}
	}
	if currentUser.IsActive != updatedUser.IsActive {
		changes["isActive"] = map[string]interface{}{
			"old": currentUser.IsActive,
			"new": updatedUser.IsActive,
		}
	}

	// Update in repository
	if err := s.userRepo.Update(ctx, updatedUser); err != nil {
		return err
	}

	// Publish user updated event if there are changes
	if len(changes) > 0 && s.publisher != nil {
		correlationID := uuid.New().String()
		userUpdatedEvent := messaging.UserUpdatedEvent{
			UserID:    updatedUser.ID.String(),
			Email:     updatedUser.Email,
			FirstName: updatedUser.Username,
			LastName:  "",
			IsActive:  updatedUser.IsActive,
			UpdatedAt: time.Now().UTC().Format(time.RFC3339),
			Changes:   changes,
		}

		if err := s.publisher.PublishUserUpdated(userUpdatedEvent, correlationID, updatedUser.ID.String()); err != nil {
			// Log error but don't fail the user update
		}
	}

	return nil
}

// DeleteUser removes a user from the system
func (s *UserService) DeleteUser(ctx context.Context, userID string) error {
	if userID == "" {
		return errors.New("user ID is required")
	}

	return s.userRepo.Delete(ctx, userID)
}
