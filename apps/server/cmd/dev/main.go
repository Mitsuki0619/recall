package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Create Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Initialize your application routes
	setupRoutes(e)

	// Start server
	log.Println("Starting server on :8080")
	e.Logger.Fatal(e.Start(":8080"))
}

func setupRoutes(e *echo.Echo) {
	// Health check
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "ok", "message": "Server is running!"})
	})

	// GraphQL endpoint (placeholder)
	e.POST("/graphql", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "GraphQL endpoint"})
	})

	// Add your application routes here
}
