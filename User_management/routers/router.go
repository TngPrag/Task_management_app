package routers

import (
	handlers "user_manager/logic/user_handlers"
	//userhandlers "user_manager/logic/user_handlers"
	middleware "user_manager/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault) // default

	//
	app.Get("/task_app/user_manager_service/api/v0.1/health", handlers.UserManagerServiceHealthCheck)
	app.Post("/task_app/user_manager_service/api/v0.1/login", handlers.User_login)
	user_group_routes := app.Group("/task_app/user_manager_service/api/v0.1/user/", middleware.JwtMiddleware)
	user_group_routes.Get("verify",handlers.User_authenticate)
	//user_group_routes.Get("login",handlers.User_login)
	user_group_routes.Post("signup",handlers.User_signup)
	user_group_routes.Post("notify", handlers.User_notify)
	user_group_routes.Get("read/:user_id",handlers.User_read_by_id)
	user_group_routes.Put("update", handlers.User_update_user_name_password)
	user_group_routes.Delete("remove/:user_id",handlers.User_remove_by_id)
	user_group_routes.Delete("remove_all",handlers.User_remove_by_owner_id)
	user_group_routes.Get("read_all",handlers.User_list_by_owner_id)

}
