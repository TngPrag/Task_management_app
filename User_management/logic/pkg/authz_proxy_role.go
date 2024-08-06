package pkg

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// GetUserRole sends an HTTP request to get the user's role
func GetUserRole(token string) (string, error) {
	url := "http://localhost:8980/task_app/authz_service/api/v0.1/role/read"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("failed to get user role: status code " + resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	var role string
	err = json.Unmarshal(body, &role)
	if err != nil {
		return "", fmt.Errorf("error unmarshaling response body: %w", err)
	}
	//log.Println(role)
	return role, nil
}

// AssignRole sends an HTTP request to assign a role to a user
func AssignRole(token string, userID string, role string) error {
	url := "http://localhost:8980/task_app/authz_service/api/v0.1/role/write"

	payload := struct {
		UserID string `json:"user_id"`
		Role   string `json:"role"`
	}{
		UserID: userID,
		Role:   role,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error marshaling payload: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("failed to assign role: status code " + resp.Status)
	}

	return nil
}
