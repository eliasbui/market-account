package main

import (
	"log"
	"user-service/application"
	"user-service/infrastructure"
	"user-service/web"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Database connection
	dsn := "host=localhost user=postgres password=postgres dbname=marketplace port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate tables
	// db.AutoMigrate(&domain.User{}, &domain.UserRole{})

	// Initialize repositories
	userRepo := infrastructure.NewPostgresUserRepository(db)

	// Initialize services
	userService := application.NewUserService(userRepo)

	// Initialize handlers
	userHandler := web.NewUserHandler(userService)

	// Setup routes
	router := gin.Default()

	// User routes
	userRoutes := router.Group("/api/v1/users")
	{
		userRoutes.POST("", userHandler.CreateUser)
		userRoutes.GET("/:id", userHandler.GetUser)
		userRoutes.GET("/email/:email", userHandler.GetUserByEmail)
		userRoutes.DELETE("/:id", userHandler.DeleteUser)
	}

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy", "service": "user-service"})
	})

	// Start server
	log.Println("Starting user service on port 8001")
	if err := router.Run(":8001"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
