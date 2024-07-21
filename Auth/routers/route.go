package routers

import (
	handlers "tele_auth/logic/auth_handlers"
	//userhandlers "user_manager/logic/user_handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault) // default
	app.Get("/task_app/authz_service/api/v0.1/health", handlers.HealthCHeckAuthzService)

	policy_routes := app.Group("/task_app/authz_service/api/v0.1/policy/")
	policy_routes.Post("write", handlers.Policy_write)
	policy_routes.Get("read/:sub", handlers.Policy_read_by_subject)
	policy_routes.Get("check_permission", handlers.Policy_check_permission)
	policy_routes.Delete("remove", handlers.Policy_remove)
	policy_routes.Get("list", handlers.Policy_list)

	role_routes := app.Group("/task_app/authz_service/api/v0.1/role/")
	role_routes.Post("write", handlers.Role_write)
	role_routes.Get("read", handlers.Role_read)

}
