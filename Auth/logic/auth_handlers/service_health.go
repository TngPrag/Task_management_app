package authhandlers

import "github.com/gofiber/fiber/v2"

func HealthCHeckAuthzService(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Authorization service health": "Ok"})
}
