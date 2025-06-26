package domain

import (
	"time"

	"github.com/google/uuid"
)

// User represents the core user entity in our domain
type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"-"` // Don't expose password in JSON
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserRole represents user roles in the system
type UserRole struct {
	ID     uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"user_id"`
	Role   string    `json:"role"`
}

// NewUser creates a new user entity
func NewUser(email, username, password string) *User {
	return &User{
		ID:        uuid.New(),
		Email:     email,
		Username:  username,
		Password:  password,
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// ValidateEmail validates user email format
func (u *User) ValidateEmail() bool {
	// Basic email validation logic
	return len(u.Email) > 0 && contains(u.Email, "@")
}

// ValidateUsername validates username requirements
func (u *User) ValidateUsername() bool {
	return len(u.Username) >= 3 && len(u.Username) <= 50
}

// contains is a helper function
func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
