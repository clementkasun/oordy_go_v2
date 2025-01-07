package controllers

import (
	"fiber_app/database"
	"fiber_app/models"
	"fiber_app/resources"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// IUserController defines the interface for user-related actions
type IUserController interface {
	CreateUser(c *fiber.Ctx) error
	GetAllUsers(c *fiber.Ctx) error
	LoginPage(c *fiber.Ctx) error
	RegisterPage(c *fiber.Ctx) error
}

// UserController implements IUserController
type UserController struct {
	Validator    *validator.Validate
	UserResource *resources.UserResource
}

// NewUserController initializes and returns a UserController
func NewUserController() IUserController {
	return &UserController{
		Validator:    validator.New(),
		UserResource: &resources.UserResource{},
	}
}

// LoginPage handles requests to display the login page
func (uc *UserController) LoginPage(c *fiber.Ctx) error {
	log.Printf("Login page requested by %s from %s", c.IP(), c.Method())
	if err := c.SendFile("./static/admin-lte-master/dist/pages/user/login.html"); err != nil {
		log.Printf("Error loading login page: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Unable to load Admin Login Page",
			"details": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(resources.SuccessResponse("Login page loaded successfully", nil))
}

// RegisterPage handles requests to display the register page
func (uc *UserController) RegisterPage(c *fiber.Ctx) error {
	log.Printf("Register page requested by %s from %s", c.IP(), c.Method())
	if err := c.SendFile("./static/admin-lte-master/dist/pages/user/register.html"); err != nil {
		log.Printf("Error loading register page: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Unable to load Register Page",
			"details": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(resources.SuccessResponse("Register page loaded successfully", nil))
}

// CreateUser handles creating a new user
func (uc *UserController) CreateUser(c *fiber.Ctx) error {
	var user models.User

	// Parse the JSON request body
	if err := c.BodyParser(&user); err != nil {
		log.Printf("Error parsing request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(resources.ErrorResponse("Failed to parse request body", err))
	}

	// Validate the user data
	if err := uc.Validator.Struct(user); err != nil {
		log.Printf("Validation error: %v", err)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(resources.ErrorResponse("Validation failed", err))
	}

	// Save the user to the database using GORM
	if err := database.DB.Create(&user).Error; err != nil {
		log.Printf("Database error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(resources.ErrorResponse("Failed to create user", err))
	}

	log.Printf("User Created: %+v", user)

	// Format the response using UserResource
	return c.Status(fiber.StatusCreated).JSON(resources.SuccessResponse("User created successfully", uc.UserResource.Transform(user)))
}

// GetAllUsers handles fetching all users from the database
func (uc *UserController) GetAllUsers(c *fiber.Ctx) error {
	var users []models.User

	// Fetch users from the database using GORM's Find method
	if err := database.DB.Find(&users).Error; err != nil {
		log.Printf("Database error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(resources.ErrorResponse("Failed to fetch users", err))
	}

	log.Printf("Users Retrieved: %+v", users)

	// Format the response using UserResource
	return c.Status(fiber.StatusOK).JSON(resources.SuccessResponse("Users retrieved successfully", uc.UserResource.TransformCollection(users)))
}
