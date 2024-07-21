package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEnforcerWrapper(t *testing.T) {
	modelPath := "../../config/model.conf"
	policyPath := "../../config/policy.csv"

	ew, err := NewEnforcerWrapper(modelPath, policyPath)
	assert.NoError(t, err)
	assert.NotNil(t, ew)
	assert.NotNil(t, ew.Enforcer)
}

func TestCreatePolicy(t *testing.T) {
	modelPath := "../../config/model.conf"
	policyPath := "../../config/policy.csv"

	ew, err := NewEnforcerWrapper(modelPath, policyPath)
	assert.NoError(t, err)

	policy := Policy{Sub: "alice", Obj: "data1", Act: "read"}
	err = ew.CreatePolicy(policy)
	assert.NoError(t, err)

	// Check if the policy was added
	policies, err := ew.ReadPoliciesForSubject("alice")
	assert.NoError(t, err)
	assert.Contains(t, policies, policy)
}

func TestReadPoliciesForSubject(t *testing.T) {
	modelPath := "../../config/model.conf"
	policyPath := "../../config/policy.csv"

	ew, err := NewEnforcerWrapper(modelPath, policyPath)
	assert.NoError(t, err)

	policy := Policy{Sub: "alice", Obj: "data1", Act: "read"}
	_ = ew.CreatePolicy(policy)

	policies, err := ew.ReadPoliciesForSubject("alice")
	assert.NoError(t, err)
	assert.Contains(t, policies, policy)
}

func TestGetAllPolicy(t *testing.T) {
	modelPath := "../../config/model.conf"
	policyPath := "../../config/policy.csv"

	ew, err := NewEnforcerWrapper(modelPath, policyPath)
	assert.NoError(t, err)

	policy1 := Policy{Sub: "alice", Obj: "data1", Act: "read"}
	policy2 := Policy{Sub: "bob", Obj: "data2", Act: "write"}
	_ = ew.CreatePolicy(policy1)
	_ = ew.CreatePolicy(policy2)

	policies, err := ew.GetAllPolicy()
	assert.NoError(t, err)
	assert.Contains(t, policies, policy1)
	assert.Contains(t, policies, policy2)
}

func TestUpdatePolicy(t *testing.T) {
	modelPath := "../../config/model.conf"
	policyPath := "../../config/policy.csv"

	ew, err := NewEnforcerWrapper(modelPath, policyPath)
	assert.NoError(t, err)

	oldPolicy := Policy{Sub: "alice", Obj: "data1", Act: "read"}
	newPolicy := Policy{Sub: "alice", Obj: "data1", Act: "write"}
	_ = ew.CreatePolicy(oldPolicy)

	err = ew.UpdatePolicy(oldPolicy, newPolicy)
	assert.NoError(t, err)

	// Check if the old policy was removed and the new policy was added
	policies, err := ew.ReadPoliciesForSubject("alice")
	assert.NoError(t, err)
	assert.NotContains(t, policies, oldPolicy)
	assert.Contains(t, policies, newPolicy)
}

func TestDeletePolicy(t *testing.T) {
	modelPath := "../../config/model.conf"
	policyPath := "../../config/policy.csv"

	ew, err := NewEnforcerWrapper(modelPath, policyPath)
	assert.NoError(t, err)

	policy := Policy{Sub: "alice", Obj: "data1", Act: "read"}
	_ = ew.CreatePolicy(policy)

	err = ew.DeletePolicy(policy)
	assert.NoError(t, err)

	// Check if the policy was removed
	policies, err := ew.ReadPoliciesForSubject("alice")
	assert.NoError(t, err)
	assert.NotContains(t, policies, policy)
}

func TestCheckPermission(t *testing.T) {
	modelPath := "../../config/model.conf"
	policyPath := "../../config/policy.csv"

	ew, err := NewEnforcerWrapper(modelPath, policyPath)
	assert.NoError(t, err)

	policy := Policy{Sub: "alice", Obj: "data1", Act: "read"}
	_ = ew.CreatePolicy(policy)

	ok, err := ew.CheckPermission("alice", "data1", "read")
	assert.NoError(t, err)
	assert.True(t, ok)

	ok, err = ew.CheckPermission("alice", "data1", "write")
	assert.NoError(t, err)
	assert.False(t, ok)
}
