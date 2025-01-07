package controllers

import "github.com/gofiber/fiber/v2"

// AdminController handles admin-related requests
var AdminController = &struct {
	AdminDashboard  func(c *fiber.Ctx) error
	AdminDashboardV2 func(c *fiber.Ctx) error
	AdminDashboardV3 func(c *fiber.Ctx) error
}{
	AdminDashboard: func(c *fiber.Ctx) error {
		if err := c.SendFile("./static/admin-lte-master/dist/pages/index.html"); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Unable to load Admin Dashboard V1",
				"details": err.Error(),
			})
		}
		return nil
	},
	AdminDashboardV2: func(c *fiber.Ctx) error {
		if err := c.SendFile("./static/admin-lte-master/dist/pages/index2.html"); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Unable to load Admin Dashboard V2",
				"details": err.Error(),
			})
		}
		return nil
	},
	AdminDashboardV3: func(c *fiber.Ctx) error {
		if err := c.SendFile("./static/admin-lte-master/dist/pages/index3.html"); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Unable to load Admin Dashboard V3",
				"details": err.Error(),
			})
		}
		return nil
	},
}
