package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRole(t *testing.T) {
	modelPath := "../../config/model.conf"
	policyPath := "../../config/policy.csv"

	ew, err := NewEnforcerWrapper(modelPath, policyPath)
	assert.NoError(t, err)

	role := Role{User: "alice", Role: "admin"}
	err = ew.CreateRole(role)
	assert.NoError(t, err)

	// Check if the role was added
	roles, err := ew.GetAllRoles()
	assert.NoError(t, err)
	assert.Contains(t, roles, role)
}

func TestGetRole(t *testing.T) {
	modelPath := "../../config/model.conf"
	policyPath := "../../config/policy.csv"

	ew, err := NewEnforcerWrapper(modelPath, policyPath)
	assert.NoError(t, err)

	role := Role{User: "alice", Role: "admin"}
	_ = ew.CreateRole(role)

	retrievedRole, err := ew.GetRole("alice")
	assert.NoError(t, err)
	assert.Equal(t, role.Role, retrievedRole)
}

func TestGetAllRoles(t *testing.T) {
	modelPath := "../../config/model.conf"
	policyPath := "../../config/policy.csv"

	ew, err := NewEnforcerWrapper(modelPath, policyPath)
	assert.NoError(t, err)

	role1 := Role{User: "alice", Role: "admin"}
	role2 := Role{User: "bob", Role: "user"}
	_ = ew.CreateRole(role1)
	_ = ew.CreateRole(role2)

	roles, err := ew.GetAllRoles()
	assert.NoError(t, err)
	assert.Contains(t, roles, role1)
	assert.Contains(t, roles, role2)
}

func TestUpdateRole(t *testing.T) {
	modelPath := "../../config/model.conf"
	policyPath := "../../config/policy.csv"

	ew, err := NewEnforcerWrapper(modelPath, policyPath)
	assert.NoError(t, err)

	oldRole := Role{User: "alice", Role: "user"}
	newRole := Role{User: "alice", Role: "admin"}
	_ = ew.CreateRole(oldRole)

	err = ew.UpdateRole(oldRole, newRole)
	assert.NoError(t, err)

	// Check if the old role was removed and the new role was added
	roles, err := ew.GetAllRoles()
	assert.NoError(t, err)
	assert.NotContains(t, roles, oldRole)
	assert.Contains(t, roles, newRole)
}

func TestDeleteRole(t *testing.T) {
	modelPath := "../../config/model.conf"
	policyPath := "../../config/policy.csv"

	ew, err := NewEnforcerWrapper(modelPath, policyPath)
	assert.NoError(t, err)

	role := Role{User: "alice", Role: "admin"}
	_ = ew.CreateRole(role)

	err = ew.DeleteRole(role)
	assert.NoError(t, err)

	// Check if the role was removed
	roles, err := ew.GetAllRoles()
	assert.NoError(t, err)
	assert.NotContains(t, roles, role)
}
