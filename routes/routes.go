package routes

import (
	"fiber_app/controllers" // Your controllers package
	"fiber_app/middlewares" // Your middlewares package

	"github.com/gofiber/fiber/v2"
)

// RegisterRoutes registers all the application routes
func RegisterRoutes(app *fiber.App) {
	// Admin routes
	registerAdminRoutes(app)

	// Initialize the user controller
	userController := controllers.NewUserController() // Initialize userController

	// User routes (uses interface-based controller)
	RegisterUserRoutes(app, userController) // Pass the userController instance
}

// registerAdminRoutes defines routes related to admin
func registerAdminRoutes(app *fiber.App) {
	adminGroup := app.Group("/admin") // Group all admin routes

	// Add middleware for admin routes (e.g., authentication)
	adminGroup.Use(middlewares.AuthMiddleware)

	// Define routes and map to controller methods
	adminGroup.Get("/", controllers.AdminController.AdminDashboard)
	adminGroup.Get("/v2", controllers.AdminController.AdminDashboardV2)
	adminGroup.Get("/v3", controllers.AdminController.AdminDashboardV3)
}

// RegisterUserRoutes registers user-specific routes using IUserController interface
func RegisterUserRoutes(app *fiber.App, userController controllers.IUserController) {
	userGroup := app.Group("/users")

	userGroup.Use(middlewares.AuthMiddleware)

	// Using the IUserController interface methods
	userGroup.Post("/createUser", userController.CreateUser)  // Create user
	userGroup.Get("/getAllUsers", userController.GetAllUsers) // Get all users

	// 	// Public routes (no middleware applied)
	userGroup.Get("/login", controllers.NewUserController().LoginPage)       // Add this method in UserController
	userGroup.Get("/register", controllers.NewUserController().RegisterPage) // Add this method in UserController
}
