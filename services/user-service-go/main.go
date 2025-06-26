package main

import (
	"log"
	"os"
	"user-service/application"
	_ "user-service/docs" // Import generated docs
	"user-service/infrastructure"
	"user-service/infrastructure/messaging"
	"user-service/web"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Set Gin mode from environment or default to debug
	if mode := os.Getenv("GIN_MODE"); mode != "" {
		gin.SetMode(mode)
	}

	// Database connection - use environment variables in production
	dsn := getDBConnectionString()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate tables
	// Uncomment when ready: db.AutoMigrate(&domain.User{})

	// Initialize RabbitMQ connection
	rabbitMQURL := getRabbitMQConnectionString()
	rabbitMQClient := messaging.NewRabbitMQClient(rabbitMQURL)

	if err := rabbitMQClient.Connect(); err != nil {
		log.Printf("Failed to connect to RabbitMQ: %v. Continuing without message publishing.", err)
		// Continue without RabbitMQ for graceful degradation
		rabbitMQClient = nil
	} else {
		defer rabbitMQClient.Close()
		log.Println("Successfully connected to RabbitMQ")
	}

	// Initialize repositories
	userRepo := infrastructure.NewPostgresUserRepository(db)

	// Initialize services with message publisher
	userService := application.NewUserService(userRepo, rabbitMQClient)

	// Initialize enhanced REST handlers
	userHandler := web.NewUserRESTHandler(userService)

	// Setup Gin router with middleware
	router := setupRouter(userHandler)

	// Start server
	port := getPort()
	log.Printf("Starting user service on port %s", port)
	log.Printf("Swagger documentation available at: http://localhost%s/swagger/index.html", port)

	if err := router.Run(port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

// setupRouter configures the Gin router with middleware and routes
func setupRouter(userHandler *web.UserRESTHandler) *gin.Engine {
	router := gin.New()

	// Add middleware in correct order
	router.Use(web.RequestLoggingMiddleware()) // Logging first
	router.Use(gin.Recovery())                 // Recovery middleware
	router.Use(web.CORSMiddleware())           // CORS handling
	router.Use(web.SecurityMiddleware())       // Security headers
	router.Use(web.RateLimitMiddleware())      // Rate limiting

	// Swagger documentation endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Health check endpoint (no auth required)
	router.GET("/health", userHandler.HealthCheck)

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// Public user routes (no authentication required)
		userRoutes := v1.Group("/users")
		{
			userRoutes.POST("", userHandler.CreateUser)                   // Create user
			userRoutes.POST("/validate", userHandler.ValidateCredentials) // Login/validate
		}

		// Protected user routes (authentication required)
		protectedUserRoutes := v1.Group("/users")
		protectedUserRoutes.Use(web.JWTAuthMiddleware())
		{
			protectedUserRoutes.GET("/:id", userHandler.GetUser)                 // Get user by ID
			protectedUserRoutes.GET("/email/:email", userHandler.GetUserByEmail) // Get user by email
			protectedUserRoutes.PUT("/:id", userHandler.UpdateUser)              // Update user
			protectedUserRoutes.DELETE("/:id", userHandler.DeleteUser)           // Delete user
		}
	}

	// Add a route for generating JWT tokens (for testing)
	if gin.Mode() == gin.DebugMode {
		router.POST("/debug/generate-token", generateTestToken)
	}

	return router
}

// generateTestToken creates a test JWT token for development/testing
func generateTestToken(c *gin.Context) {
	var req struct {
		UserID   string `json:"user_id" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Username string `json:"username" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, web.ErrorResponse{
			Error:   "validation_error",
			Message: "Invalid request data",
			Code:    "INVALID_INPUT",
		})
		return
	}

	token, err := web.GenerateJWT(req.UserID, req.Email, req.Username)
	if err != nil {
		c.JSON(500, web.ErrorResponse{
			Error:   "token_generation_failed",
			Message: err.Error(),
			Code:    "TOKEN_ERROR",
		})
		return
	}

	c.JSON(200, web.SuccessResponse{
		Message: "Token generated successfully",
		Data: map[string]interface{}{
			"token": token,
			"type":  "Bearer",
		},
	})
}

// getDBConnectionString returns database connection string
func getDBConnectionString() string {
	// In production, use environment variables
	if dsn := os.Getenv("DATABASE_URL"); dsn != "" {
		return dsn
	}

	// Development default
	return "host=localhost user=postgres password=postgres dbname=marketplace port=5432 sslmode=disable"
}

// getRabbitMQConnectionString returns RabbitMQ connection string
func getRabbitMQConnectionString() string {
	// Check for environment variable first
	if url := os.Getenv("RABBITMQ_URL"); url != "" {
		return url
	}

	// Development default
	return "amqp://admin:admin123@localhost:5672/"
}

// getPort returns the port to run the server on
func getPort() string {
	if port := os.Getenv("PORT"); port != "" {
		return ":" + port
	}
	return ":8001"
}
