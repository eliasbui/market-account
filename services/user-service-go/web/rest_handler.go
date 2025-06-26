package web

import (
	"net/http"
	"user-service/application"
	"user-service/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UserRESTHandler handles REST API requests for user operations
type UserRESTHandler struct {
	userService *application.UserService
}

// NewUserRESTHandler creates a new REST handler for users
func NewUserRESTHandler(userService *application.UserService) *UserRESTHandler {
	return &UserRESTHandler{
		userService: userService,
	}
}

// API Request/Response DTOs
type CreateUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=6"`
}

type UpdateUserRequest struct {
	Email    *string `json:"email,omitempty" binding:"omitempty,email"`
	Username *string `json:"username,omitempty" binding:"omitempty,min=3,max=50"`
	IsActive *bool   `json:"is_active,omitempty"`
}

type UserResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	IsActive  bool   `json:"is_active"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UsersListResponse struct {
	Users []UserResponse `json:"users"`
	Total int            `json:"total"`
	Page  int            `json:"page"`
	Limit int            `json:"limit"`
}

type ErrorResponse struct {
	Error   string                 `json:"error"`
	Message string                 `json:"message"`
	Details map[string]interface{} `json:"details,omitempty"`
	Code    string                 `json:"code"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// CreateUser handles POST /users
func (h *UserRESTHandler) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "validation_error",
			Message: "Invalid request data",
			Details: map[string]interface{}{"validation": err.Error()},
			Code:    "INVALID_INPUT",
		})
		return
	}

	user, err := h.userService.CreateUser(c.Request.Context(), req.Email, req.Username, req.Password)
	if err != nil {
		status := http.StatusInternalServerError
		code := "CREATION_FAILED"

		// Handle specific error cases
		if err.Error() == "user with this email already exists" {
			status = http.StatusConflict
			code = "EMAIL_EXISTS"
		}

		c.JSON(status, ErrorResponse{
			Error:   "creation_error",
			Message: err.Error(),
			Code:    code,
		})
		return
	}

	userResp := convertDomainUserToResponse(user)
	c.JSON(http.StatusCreated, SuccessResponse{
		Message: "User created successfully",
		Data:    userResp,
	})
}

// GetUser handles GET /users/:id
func (h *UserRESTHandler) GetUser(c *gin.Context) {
	userID := c.Param("id")

	// Validate UUID format
	if _, err := uuid.Parse(userID); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "validation_error",
			Message: "Invalid user ID format",
			Code:    "INVALID_USER_ID",
		})
		return
	}

	user, err := h.userService.GetUserByID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{
			Error:   "not_found",
			Message: "User not found",
			Code:    "USER_NOT_FOUND",
		})
		return
	}

	userResp := convertDomainUserToResponse(user)
	c.JSON(http.StatusOK, SuccessResponse{
		Message: "User retrieved successfully",
		Data:    userResp,
	})
}

// GetUserByEmail handles GET /users/email/:email
func (h *UserRESTHandler) GetUserByEmail(c *gin.Context) {
	email := c.Param("email")

	if email == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "validation_error",
			Message: "Email is required",
			Code:    "INVALID_EMAIL",
		})
		return
	}

	user, err := h.userService.GetUserByEmail(c.Request.Context(), email)
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{
			Error:   "not_found",
			Message: "User not found",
			Code:    "USER_NOT_FOUND",
		})
		return
	}

	userResp := convertDomainUserToResponse(user)
	c.JSON(http.StatusOK, SuccessResponse{
		Message: "User retrieved successfully",
		Data:    userResp,
	})
}

// UpdateUser handles PUT /users/:id
func (h *UserRESTHandler) UpdateUser(c *gin.Context) {
	userID := c.Param("id")

	// Validate UUID format
	if _, err := uuid.Parse(userID); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "validation_error",
			Message: "Invalid user ID format",
			Code:    "INVALID_USER_ID",
		})
		return
	}

	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "validation_error",
			Message: "Invalid request data",
			Details: map[string]interface{}{"validation": err.Error()},
			Code:    "INVALID_INPUT",
		})
		return
	}

	// Get existing user
	user, err := h.userService.GetUserByID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{
			Error:   "not_found",
			Message: "User not found",
			Code:    "USER_NOT_FOUND",
		})
		return
	}

	// Update fields if provided
	if req.Email != nil {
		user.Email = *req.Email
	}
	if req.Username != nil {
		user.Username = *req.Username
	}
	if req.IsActive != nil {
		user.IsActive = *req.IsActive
	}

	// Update user
	if err := h.userService.UpdateUser(c.Request.Context(), user); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "update_error",
			Message: err.Error(),
			Code:    "UPDATE_FAILED",
		})
		return
	}

	userResp := convertDomainUserToResponse(user)
	c.JSON(http.StatusOK, SuccessResponse{
		Message: "User updated successfully",
		Data:    userResp,
	})
}

// DeleteUser handles DELETE /users/:id
func (h *UserRESTHandler) DeleteUser(c *gin.Context) {
	userID := c.Param("id")

	// Validate UUID format
	if _, err := uuid.Parse(userID); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "validation_error",
			Message: "Invalid user ID format",
			Code:    "INVALID_USER_ID",
		})
		return
	}

	if err := h.userService.DeleteUser(c.Request.Context(), userID); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "deletion_error",
			Message: err.Error(),
			Code:    "DELETE_FAILED",
		})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Message: "User deleted successfully",
	})
}

// ValidateCredentials handles POST /users/validate
func (h *UserRESTHandler) ValidateCredentials(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "validation_error",
			Message: "Invalid request data",
			Details: map[string]interface{}{"validation": err.Error()},
			Code:    "INVALID_INPUT",
		})
		return
	}

	isValid, user, err := h.userService.ValidateCredentials(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "validation_error",
			Message: err.Error(),
			Code:    "VALIDATION_FAILED",
		})
		return
	}

	response := map[string]interface{}{
		"is_valid": isValid,
	}

	if isValid && user != nil {
		response["user"] = convertDomainUserToResponse(user)
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Message: "Credentials validated",
		Data:    response,
	})
}

// Health check endpoint
func (h *UserRESTHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, SuccessResponse{
		Message: "User service is healthy",
		Data: map[string]interface{}{
			"service": "user-service",
			"status":  "up",
			"version": "1.0.0",
		},
	})
}

// Helper function to convert domain user to API response
func convertDomainUserToResponse(user *domain.User) UserResponse {
	return UserResponse{
		ID:        user.ID.String(),
		Email:     user.Email,
		Username:  user.Username,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}
}
