package main

import (
	"fmt"
	"log"

	v1Handler "example.com/golang-todo/internal/api/v1/handler"
	"example.com/golang-todo/internal/app/http/middleware"
	"example.com/golang-todo/internal/app/http/routes"
	"example.com/golang-todo/internal/config"
	"example.com/golang-todo/internal/repository"
	"example.com/golang-todo/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load config values
	config.Load()

	// Initialize repositories
	userRepo := repository.NewUserRepository()

	// Initialize services
	userService := service.NewUserService(userRepo)

	// Initialize handlers
	userHandler := v1Handler.NewUserHandler(userService)

	// Initialize routes
	userRoute := routes.NewUserRoute(userHandler)

	r := gin.Default()

	// Register routes
	routes.RegisterRoutes(r, userRoute)

	// Setup rate limit middleware
	go middleware.CleanupRateLimiters()
	r.Use(middleware.RateLimitByIp())

	// Load configs
	config, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load server configs!")
	}

	// Start the server
	if err := r.Run(fmt.Sprintf(":%s", config.Server.Port)); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
