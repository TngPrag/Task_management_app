package userhandlers

import "github.com/gofiber/fiber/v2"

// UserManagerServiceHealthCheck godoc
// @Summary Check the health of the user management service
// @Description Get the health status of the user management service
// @Tags health
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func UserManagerServiceHealthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Health check user management service":"Ok"})
}
