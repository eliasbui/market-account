package application

import (
	"context"
	"errors"
	"user-service/domain"
)

// UserRepository defines the interface for user data access
type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	GetByID(ctx context.Context, id string) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, id string) error
}

// UserService implements the application use cases for user management
type UserService struct {
	userRepo UserRepository
}

func (s *UserService) GetUser(ctx context.Context, id string) (*domain.User, error) {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// NewUserService creates a new instance of UserService
func NewUserService(userRepo UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
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
func (s *UserService) UpdateUser(ctx context.Context, user *domain.User) error {
	if user == nil {
		return errors.New("user is required")
	}

	return s.userRepo.Update(ctx, user)
}

// DeleteUser removes a user from the system
func (s *UserService) DeleteUser(ctx context.Context, userID string) error {
	if userID == "" {
		return errors.New("user ID is required")
	}

	return s.userRepo.Delete(ctx, userID)
}
