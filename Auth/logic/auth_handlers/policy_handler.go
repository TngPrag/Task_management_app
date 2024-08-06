package authhandlers

import (
	"encoding/json"
	"tele_auth/logic/core"
	"tele_auth/logic/dto"
	"tele_auth/middlewares"

	"github.com/gofiber/fiber/v2"
)

// super-admin

// PolicyWrite godoc
// @Summary Create a new policy
// @Description Create a new RBAC policy
// @Tags Policy
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param policy body dto.CreatePolicyDto true "Policy to be created"
// @Success 200 {object} core.Policy
// @Failure 400 {object} map[string]string
// @Security BearerAuth
// @Router /policy/write [post]
func Policy_write(c *fiber.Ctx) error {
	policy_write_dto := new(dto.CreatePolicyDto)
	if err := c.BodyParser(policy_write_dto); err != nil {
		return c.Status(400).JSON(fiber.Map{"Invalid input, create user": err.Error()})
	}
	if err := policy_write_dto.ValidateCreatePolicyDto(); err != nil {
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
	check_permission, _ := enforceWrapper.CheckPermission(user_role, "task_app/authz_service/api/v0.1/policy", "POST")
	if check_permission {
		policy := core.Policy{Sub: policy_write_dto.Subject, Obj: policy_write_dto.Object, Act: policy_write_dto.Action}
		if err := enforceWrapper.CreatePolicy(core.Policy(policy)); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error: while creating policy": err.Error()})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"Status": "Policy created succefully!"})
	}
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"User Unauthorized!": err})

}

// super-admin/

// PolicyReadBySubject godoc
// @Summary Get policy by subject
// @Description Get RBAC policy by subject
// @Tags Policy
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param sub path string true "Subject"
// @Success 200 {object} core.Policy
// @Failure 404 {object} map[string]string
// @Security BearerAuth
// @Router /policy/read/{sub} [get]
func Policy_read_by_subject(c *fiber.Ctx) error {
	subject := c.Params("sub")
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
	// read policies of the user
	//policy := new(core.Policy)
	//list_user_policies := []core.Policy{}
	user_role, err := enforceWrapper.GetRole(user_id)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"User Unauthorized!": err.Error()})
	}
	check_permission, _ := enforceWrapper.CheckPermission(user_role, "task_app/authz_service/api/v0.1/policy", "GET")
	if check_permission {
		user_policies, err := enforceWrapper.ReadPoliciesForSubject(subject)
		if err != nil {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"No policies defined for the subject": err})
		}
		user_policy, _ := json.Marshal(user_policies)
		return c.Status(fiber.StatusOK).JSON(user_policy)
	}
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "user has no permission"})
}

// PolicyRemove godoc
// @Summary Remove a policy
// @Description Remove an existing RBAC policy
// @Tags Policy
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param policy body dto.RemovePolicyDto true "Policy to be removed"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Security BearerAuth
// @Router /policy/remove [delete]
func Policy_remove(c *fiber.Ctx) error {
	rem_policy_dto := new(dto.RemovePolicyDto)
	if err := c.BodyParser(rem_policy_dto); err != nil {
		return c.Status(400).JSON(fiber.Map{"Invalid input, remove policy": err.Error()})
	}
	if err := rem_policy_dto.ValidateRemovePolicyDto(); err != nil {
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
	check_permission, _ := enforceWrapper.CheckPermission(user_role, "task_app/authz_service/api/v0.1/policy", "DELETE")
	if check_permission {
		policy := core.Policy{Sub: rem_policy_dto.Subject, Obj: rem_policy_dto.Object, Act: rem_policy_dto.Action}
		if err := enforceWrapper.DeletePolicy(core.Policy(policy)); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error: while removing policy": err.Error()})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"Status": "Policy deleted succefully!"})
	}
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"User Unauthorized!": err})
}

// PolicyCheckPermission godoc
// @Summary Check permission
// @Description Check if a subject has a specific permission
// @Tags Policy
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param check body dto.CheckPolicyPermissionDto true "Permission check request"
// @Success 200 {object} core.Policy
// @Failure 400 {object} map[string]string
// @Security BearerAuth
// @Router /policy/check_permission [post]
func Policy_check_permission(c *fiber.Ctx) error {
	check_pol_perm_dto := new(dto.CheckPolicyPermissionDto)
	if err := c.BodyParser(check_pol_perm_dto); err != nil {
		return c.Status(400).JSON(fiber.Map{"Invalid input, remove policy": err.Error()})
	}
	if err := check_pol_perm_dto.ValidateCheckPolicyPermissionDto(); err != nil {
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

	check_permission, _ := enforceWrapper.CheckPermission(user_role, "task_app/authz_service/api/v0.1/policy", "Verify")
	if check_permission {
		//policy := core.Policy{Sub: check_pol_perm_dto.Subject, Obj: check_pol_perm_dto.Object, Act: check_pol_perm_dto.Action}
		status, err := enforceWrapper.CheckPermission(check_pol_perm_dto.Subject, check_pol_perm_dto.Object, check_pol_perm_dto.Action)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error: while removing policy": err.Error()})
		}
		return c.Status(fiber.StatusOK).JSON(status)
	}
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"User Unauthorized!": err})
}

// PolicyList godoc
// @Summary List all policies
// @Description List all RBAC policies
// @Tags Policy
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Success 200 {array} core.Policy
// @Security BearerAuth
// @Router /policy/list [get]
func Policy_list(c *fiber.Ctx) error {
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
	check_permission, _ := enforceWrapper.CheckPermission(user_role, "task_app/authz_service/api/v0.1/policy", "GET")
	if check_permission {
		//policy := core.Policy{Sub: che, Obj: check_pol_perm_dto.Object, Act: check_pol_perm_dto.Action}
		policies, err := enforceWrapper.GetAllPolicy()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error: while removing policy": err.Error()})
		}
		return c.Status(fiber.StatusOK).JSON(policies)
	}
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"User Unauthorized!": err})
}
