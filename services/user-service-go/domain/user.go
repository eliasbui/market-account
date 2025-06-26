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

// ValidatePassword checks if the provided password matches the user's password
func (u *User) ValidatePassword(password string) bool {
	// In a real implementation, you would hash the input password
	// and compare it with the stored hashed password
	// For now, we'll do a simple string comparison (NOT for production!)
	return u.Password == password
}

// ValidateEmail checks if the user's email is in a valid format
func (u *User) ValidateEmail() bool {
	// Basic email validation - in production use a proper email validation library
	return len(u.Email) > 0 && len(u.Email) <= 255 &&
		u.Email != "" &&
		len(u.Email) > 3 // Very basic check
}

// ValidateUsername checks if the username meets requirements
func (u *User) ValidateUsername() bool {
	return len(u.Username) >= 3 && len(u.Username) <= 50
}

// UpdatePassword updates the user's password
func (u *User) UpdatePassword(newPassword string) {
	// In production, hash the password before storing
	u.Password = newPassword
	u.UpdatedAt = time.Now()
}

// Activate activates the user account
func (u *User) Activate() {
	u.IsActive = true
	u.UpdatedAt = time.Now()
}

// Deactivate deactivates the user account
func (u *User) Deactivate() {
	u.IsActive = false
	u.UpdatedAt = time.Now()
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
