package core

import (
	"fmt"
)

type Role struct {
	User string
	Role string
}

func (e *EnforceWrapper) CreateRole(role Role) error {
	if added, err := e.Enforcer.AddGroupingPolicy(role.User, role.Role); err != nil {
		return fmt.Errorf("error adding role: %w", err)
	} else if !added {
		return fmt.Errorf("role already exists")
	}
	return e.Enforcer.SavePolicy()
}

func (e *EnforceWrapper) GetRole(user_id string)(string, error){
	roles, err := e.Enforcer.GetGroupingPolicy()
	if err != nil {
		return "", fmt.Errorf("no roles found!:%w", err)
	}
	var result []Role
	for _, r := range roles {
		result = append(result, Role{User: r[0],Role: r[1]})
	}
	var role string
	for _, user_role := range result {
		if user_role.User == user_id{
			role = user_role.Role
			break
		}
	}
	if role == ""{
		return "", fmt.Errorf("User role not found")
	} else {
		return role, nil
	}
}

func (e *EnforceWrapper) GetAllRoles() ([]Role, error) {
	roles, err := e.Enforcer.GetGroupingPolicy()
	if err != nil {
		return nil, fmt.Errorf("no roles found!:%w", err)
	}
	var result []Role
	for _, r := range roles {
		result = append(result, Role{User: r[0],Role: r[1]})
	}
	return result, nil
}

func (e *EnforceWrapper) UpdateRole(oldRole, newRole Role) error {
	if removed, err := e.Enforcer.RemoveGroupingPolicy(oldRole.User,oldRole.Role); err!= nil{
		return fmt.Errorf("error removing roles: %w", err)
	} else if !removed {
		return fmt.Errorf("error role to update not found")
	}
	if added, err := e.Enforcer.AddGroupingPolicy(newRole.User,newRole.Role); err != nil {
		return fmt.Errorf("error adding new roles!%w", err)
	} else if !added {
		return fmt.Errorf("new roles already exists")
	}
	return e.Enforcer.SavePolicy()
}

func (e *EnforceWrapper) DeleteRole(role Role) error {
	if removed, err := e.Enforcer.RemoveGroupingPolicy(role.User,role.Role); err != nil {
		return fmt.Errorf("error removing role:%w", err)
	} else if !removed {
		return fmt.Errorf("role not found")
	}
	return e.Enforcer.SavePolicy()
}

// func dmain(){
// 	enforceWrapper, err:= NewEnforcerWrapper("model.conf","postgresql://username:password@localhost:5432/dbname")
// 	if err!= nil{
// 		fmt.Println("Error:",err)
// 	}
// 	role := Role{User: "Tsegay",Role: "super-admin"}
// 	if err:= enforceWrapper.CreateRole(role);err!= nil{
// 		fmt.Println("CreatedRole error",err)
// 	}
// }