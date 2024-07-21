package userhandlers

import "github.com/gofiber/fiber/v2"

func UserManagerServiceHealthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Health check user management service":"Ok"})
}
