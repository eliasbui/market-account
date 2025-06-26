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

// NewUserService creates a new instance of UserService
func NewUserService(userRepo UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// CreateUser creates a new user in the system
func (s *UserService) CreateUser(ctx context.Context, email, username, password string) (*domain.User, error) {
	// Create new user entity
	user := domain.NewUser(email, username, password)

	// Validate user data
	if !user.ValidateEmail() {
		return nil, errors.New("invalid email format")
	}

	if !user.ValidateUsername() {
		return nil, errors.New("invalid username: must be between 3 and 50 characters")
	}

	// Check if user already exists
	existingUser, _ := s.userRepo.GetByEmail(ctx, email)
	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	// Create user in repository
	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

// GetUser retrieves a user by ID
func (s *UserService) GetUser(ctx context.Context, id string) (*domain.User, error) {
	return s.userRepo.GetByID(ctx, id)
}

// GetUserByEmail retrieves a user by email
func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	return s.userRepo.GetByEmail(ctx, email)
}

// UpdateUser updates an existing user
func (s *UserService) UpdateUser(ctx context.Context, user *domain.User) error {
	if !user.ValidateEmail() {
		return errors.New("invalid email format")
	}

	if !user.ValidateUsername() {
		return errors.New("invalid username: must be between 3 and 50 characters")
	}

	return s.userRepo.Update(ctx, user)
}

// DeleteUser deletes a user by ID
func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	return s.userRepo.Delete(ctx, id)
}
