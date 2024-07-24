package authhandlers

import (
	"tele_auth/logic/core"
	"tele_auth/logic/dto"
	"tele_auth/middlewares"

	"github.com/gofiber/fiber/v2"
)

// should be allowed to super-admin/admin

// RoleWrite godoc
// @Summary Create a new role
// @Description Create a new role for RBAC
// @Tags Role
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param role body dto.CreateRoleDto true "Role to be created"
// @Success 200 {object} core.Role
// @Failure 400 {object} map[string]string
// @Security BearerAuth
// @Router /role/write [post]
func Role_write(c *fiber.Ctx) error {
	role_write_dto := new(dto.CreateRoleDto)
	if err := c.BodyParser(role_write_dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Invalid input, create user": err.Error()})
	}
	if err := role_write_dto.ValidateCreateRoleDto(); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"Error": err.Error()})
	}
	// authenticate the user
	authHeader := c.Get("Authorization")
	user_id, _, err := middlewares.Authenticate_user(authHeader)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"User is unknown": err.Error()})
	}
	// check authorization of the user_id by reading its role and its defined policies
	enforceWrapper, err := core.NewEnforcerWrapper("config/model.conf", "config/policy.csv")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}
	user_role, err := enforceWrapper.GetRole(user_id)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"User Unauthorized!": err.Error()})
	}
	check_permission, _ := enforceWrapper.CheckPermission(user_role, "task_app/authz_service/api/v0.1/role", "POST")
	if check_permission {
		if user_role == "super-admin" {
			if role_write_dto.Role == "admin" {
				role := core.Role{User: role_write_dto.UserID, Role: role_write_dto.Role}
				if err := enforceWrapper.CreateRole(core.Role(role)); err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error: while creating role": err.Error()})
				}
				return c.Status(fiber.StatusOK).JSON(fiber.Map{"Status": "role created succefully!"})
			} else {
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"Super-admin has no previlage to assign a role to a user": err})
			}
		} else if user_role == "admin" {
			if role_write_dto.Role == "user" {
				role := core.Role{User: role_write_dto.UserID, Role: role_write_dto.Role}
				if err := enforceWrapper.CreateRole(core.Role(role)); err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error: while creating role": err.Error()})
				}
				return c.Status(fiber.StatusOK).JSON(fiber.Map{"Status": "role created succefully!"})
			} else {
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"admin has no previlage to assign a role to other that user": err})
			}
		}

	}
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"User Unauthorized!": err})
}

// should be allowed to super-admin/admin/user

// RoleRead godoc
// @Summary Get all roles
// @Description Get all roles in the system
// @Tags Role
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Success 200 {array} core.Role
// @Security BearerAuth
// @Router /role/read [get]
func Role_read(c *fiber.Ctx) error {
	//user_id := c.Params("user_id")
	// authenticate the user
	authHeader := c.Get("Authorization")
	uid, _, err := middlewares.Authenticate_user(authHeader)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"User is unknown": err.Error()})
	}
	// check authorization of the user_id by reading its role and its defined policies
	enforceWrapper, err := core.NewEnforcerWrapper("config/model.conf", "config/policy.csv")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}
	user_role, err := enforceWrapper.GetRole(uid)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"User Unauthorized!": err.Error()})
	}

	check_permission, _ := enforceWrapper.CheckPermission(user_role, "task_app/authz_service/api/v0.1/role", "GET")
	if check_permission {
		user_role, err := enforceWrapper.GetRole(uid)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"user role undefined": err.Error()})
		}
		return c.Status(fiber.StatusOK).JSON(user_role)
	}
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"User Unauthorized!": err})

}
