package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// VerifyPolicy sends an HTTP request to verify if a policy allows a certain action
func VerifyPolicy(token, sub, obj, act string) (bool, error) {
	url := "http://localhost:8980/task_app/authz_service/api/v0.1/policy/check_permission"

	payload := struct {
		Sub    string `json:"sub"`
		Object string `json:"object"`
		Action string `json:"action"`
	}{
		Sub:    sub,
		Object: obj,
		Action: act,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return false, fmt.Errorf("error marshaling payload: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return false, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("failed to verify policy: status code %s", resp.Status)
	}
	body, err1 := io.ReadAll(resp.Body)
	if err1 != nil {
		return false, fmt.Errorf("error reading response body: %w", err)
	}

	var status bool
	//log.Println("status1:", string(body))
	err = json.Unmarshal(body, &status)
	if err != nil {
		return false, fmt.Errorf("error unmarshaling response body: %w", err)
	}
	//log.Println("status:", status)
	// var response struct {
	// 	Allowed bool `json:"allowed"`
	// }
	// //log.Println(string(resp.Body))
	// err = json.NewDecoder(resp.Body).Decode(&response)
	// if err != nil {
	// 	return false, fmt.Errorf("error unmarshaling response body: %w", err)
	// }

	return status, nil
}
