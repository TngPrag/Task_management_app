package core

import (
	"fmt"

	"github.com/casbin/casbin/v2"
)

type Policy struct {
	Sub string
	Obj string
	Act string
}
type EnforceWrapper struct {
	Enforcer *casbin.Enforcer
}

func NewEnforcerWrapper(model_path, policy_path string) (*EnforceWrapper, error) {
	// Initialize a Go-pg adapter and use it in a Casbin enforcer:
	// The adapter will use the Postgres database named "casbin".
	// If it doesn't exist, the adapter will create it automatically.
	//adapter, err := pgadapter.NewAdapter("postgresql://username:password@postgres:5432/database?sslmode=disable") // Your driver and data source.
	// Alternatively, you can construct an adapter instance with *pg.Options:
	// a, _ := pgadapter.NewAdapter(&pg.Options{
	//     Database: "...",
	//     User: "...",
	//     Password: "...",
	// })

	// The adapter will use the table named "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.

	// Or you can use an existing DB by adding a second string parameter with your database name to the NewAdapter(), like this:
	// a, _ := pgadapter.NewAdapter("postgresql://username:password@postgres:5432/database?sslmode=disable", "your_database_name")
	//if err != nil {
	//	return nil, fmt.Errorf("error initializing adapter: %w", err)
	//}
	// enforcer, err := casbin.NewEnforcer(model_path, adapter)
	// if err != nil {
	// 	return nil, fmt.Errorf("error initializing enforcer: %w", err)
	// }
	enforcer, err := casbin.NewEnforcer(model_path, policy_path)
	if err != nil {
		return nil, fmt.Errorf("error initializing enforcer: %w", err)
	}
	if err := enforcer.LoadPolicy(); err != nil {
		return nil, fmt.Errorf("error loading policy: %w", err)
	}
	return &EnforceWrapper{Enforcer: enforcer}, nil
}

func (e *EnforceWrapper) CreatePolicy(policy Policy) error {
	if added, err := e.Enforcer.AddPolicy(policy.Sub, policy.Obj, policy.Act); err != nil {
		return fmt.Errorf("error adding policy: %w", err)
	} else if !added {
		return fmt.Errorf("policy already exists")
	}
	return e.Enforcer.SavePolicy()
}

func (e *EnforceWrapper) ReadPoliciesForSubject(sub string) ([]Policy, error) {
	policies, err := e.Enforcer.GetFilteredPolicy(0, sub)
	if err != nil {
		return nil, fmt.Errorf("polciy of the given subject not found:%w", err)
	}
	var result []Policy
	for _, p := range policies {
		result = append(result, Policy{Sub: p[0], Obj: p[1], Act: p[2]})
	}
	return result, nil
}
func (e *EnforceWrapper) GetAllPolicy() ([]Policy, error) {
	policies, err := e.Enforcer.GetPolicy()
	if err != nil {
		return nil, fmt.Errorf("no policies found!:%w", err)
	}
	var result []Policy
	for _, p := range policies {
		result = append(result, Policy{Sub: p[0], Obj: p[1], Act: p[2]})
	}
	return result, err
}

func (e *EnforceWrapper) UpdatePolicy(oldPolicy, newPolicy Policy) error {
	if removed, err := e.Enforcer.RemovePolicy(oldPolicy.Sub, oldPolicy.Obj, oldPolicy.Act); err != nil {
		return fmt.Errorf("error removing policy: %w", err)
	} else if !removed {
		return fmt.Errorf("error policy to update not found")
	}
	if added, err := e.Enforcer.AddPolicy(newPolicy.Sub, newPolicy.Obj, newPolicy.Act); err != nil {
		return fmt.Errorf("error adding new policy!%w", err)
	} else if !added {
		return fmt.Errorf("new policies already exists")
	}
	return e.Enforcer.SavePolicy()
}

func (e *EnforceWrapper) DeletePolicy(policy Policy) error {
	if removed, err := e.Enforcer.RemovePolicy(policy.Sub, policy.Obj, policy.Act); err != nil {
		return fmt.Errorf("error removing policy:%w", err)
	} else if !removed {
		return fmt.Errorf("polciy not found")
	}
	return e.Enforcer.SavePolicy()
}

func (e *EnforceWrapper) CheckPermission(sub, obj, act string) (bool, error) {
	ok, err := e.Enforcer.Enforce(sub, obj, act)
	if err != nil {
		return ok, fmt.Errorf("error enforcing policy:%w", err)
	}
	if ok {
		return ok, nil
	} else {
		return ok, nil
	}
}
// func Dmain() {
// 	enforceWrapper, err := NewEnforcerWrapper("model.conf", "postgresql://username:password@localhost:5432/dbname")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	}
// 	policy := Policy{Sub: "driver", Obj: "luwa/api/v0.1/driver", Act: "POST"}
// 	if err := enforceWrapper.CreatePolicy(policy); err != nil {
// 		fmt.Errorf("create policy error", err)
// 	}

// }
