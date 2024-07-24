package authhandlers

import "github.com/gofiber/fiber/v2"

// HealthCheckAuthzService godoc
// @Summary Health check for the authorization service
// @Description Returns the health status of the authorization service
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func HealthCHeckAuthzService(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Authorization service health": "Ok"})
}
