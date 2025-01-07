package main

import (
	"fiber_app/database" // Only import the database package, not anything else that imports main
	"fiber_app/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Ensure JWT_SECRET is set
	if os.Getenv("JWT_SECRET") == "" {
		log.Fatal("JWT_SECRET not set in environment variables")
	}

	// Initialize the database
	database.ConnectDatabase() // No error handling required here, as ConnectDatabase does not return an error

	// Create a new Fiber app
	app := fiber.New()

	// Serve static files from the "static" directory
	app.Static("/static", "./static")
	log.Println("Serving static files from ./static")

	// Register all the routes
	routes.RegisterRoutes(app)

	// Start the server
	serverAddress := ":3000"
	log.Printf("Starting server on %s\n", serverAddress)
	if err := app.Listen(serverAddress); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

	log.Println("Server started successfully.")
}
