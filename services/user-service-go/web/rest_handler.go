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
	Email    string `json:"email" binding:"required,email" example:"user@example.com"`
	Username string `json:"username" binding:"required,min=3,max=50" example:"johndoe"`
	Password string `json:"password" binding:"required,min=6" example:"password123"`
}

type UpdateUserRequest struct {
	Email    *string `json:"email,omitempty" binding:"omitempty,email" example:"updated@example.com"`
	Username *string `json:"username,omitempty" binding:"omitempty,min=3,max=50" example:"newusername"`
	IsActive *bool   `json:"is_active,omitempty" example:"true"`
}

type UserResponse struct {
	ID        string `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Email     string `json:"email" example:"user@example.com"`
	Username  string `json:"username" example:"johndoe"`
	IsActive  bool   `json:"is_active" example:"true"`
	CreatedAt string `json:"created_at" example:"2024-01-15T10:30:45Z"`
	UpdatedAt string `json:"updated_at" example:"2024-01-15T10:30:45Z"`
}

type UsersListResponse struct {
	Users []UserResponse `json:"users"`
	Total int            `json:"total" example:"100"`
	Page  int            `json:"page" example:"1"`
	Limit int            `json:"limit" example:"20"`
}

type ErrorResponse struct {
	Error   string                 `json:"error" example:"validation_error"`
	Message string                 `json:"message" example:"Invalid request data"`
	Details map[string]interface{} `json:"details,omitempty"`
	Code    string                 `json:"code" example:"INVALID_INPUT"`
}

type SuccessResponse struct {
	Message string      `json:"message" example:"Operation completed successfully"`
	Data    interface{} `json:"data,omitempty"`
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with email, username, and password
// @Tags users
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "User creation data"
// @Success 201 {object} SuccessResponse{data=UserResponse} "User created successfully"
// @Failure 400 {object} ErrorResponse "Invalid request data"
// @Failure 409 {object} ErrorResponse "Email already exists"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/v1/users [post]
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

// GetUser godoc
// @Summary Get user by ID
// @Description Get a single user by their UUID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID (UUID)"
// @Success 200 {object} SuccessResponse{data=UserResponse} "User retrieved successfully"
// @Failure 400 {object} ErrorResponse "Invalid user ID format"
// @Failure 404 {object} ErrorResponse "User not found"
// @Router /api/v1/users/{id} [get]
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

// GetUserByEmail godoc
// @Summary Get user by email
// @Description Get a single user by their email address
// @Tags users
// @Accept json
// @Produce json
// @Param email path string true "User email address"
// @Success 200 {object} SuccessResponse{data=UserResponse} "User retrieved successfully"
// @Failure 400 {object} ErrorResponse "Invalid email"
// @Failure 404 {object} ErrorResponse "User not found"
// @Router /api/v1/users/email/{email} [get]
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

// UpdateUser godoc
// @Summary Update user
// @Description Update an existing user's information
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID (UUID)"
// @Param user body UpdateUserRequest true "User update data"
// @Success 200 {object} SuccessResponse{data=UserResponse} "User updated successfully"
// @Failure 400 {object} ErrorResponse "Invalid request data or user ID"
// @Failure 404 {object} ErrorResponse "User not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/v1/users/{id} [put]
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

// DeleteUser godoc
// @Summary Delete user
// @Description Delete a user by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID (UUID)"
// @Success 200 {object} SuccessResponse "User deleted successfully"
// @Failure 400 {object} ErrorResponse "Invalid user ID format"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/v1/users/{id} [delete]
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

// ValidateCredentials godoc
// @Summary Validate user credentials
// @Description Validate user email and password for authentication
// @Tags users
// @Accept json
// @Produce json
// @Param credentials body object{email=string,password=string} true "User credentials"
// @Success 200 {object} SuccessResponse{data=object{is_valid=boolean,user=UserResponse}} "Credentials validated"
// @Failure 400 {object} ErrorResponse "Invalid request data"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/v1/users/validate [post]
func (h *UserRESTHandler) ValidateCredentials(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email" example:"user@example.com"`
		Password string `json:"password" binding:"required" example:"password123"`
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

// HealthCheck godoc
// @Summary Health check
// @Description Check the health status of the user service
// @Tags health
// @Produce json
// @Success 200 {object} SuccessResponse{data=object{service=string,status=string,version=string}} "Service is healthy"
// @Router /health [get]
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
