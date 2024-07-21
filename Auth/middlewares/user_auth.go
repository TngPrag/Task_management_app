package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// AuthenticateUser authenticates the user based on the provided authorization header.
func Authenticate_user(authHeader string) (string, string, error) {
	if strings.HasPrefix(authHeader, "Bearer ") && authHeader != "" {
		userToken := strings.TrimPrefix(authHeader, "Bearer ")
		userID, _, _, _, err := ViewUserProfile(userToken)
		if err != nil {
			return "", "", err
		}
		return userID, userToken, nil
	} else {
		return "", "", fmt.Errorf("invalid authorization header")
	}
}

// ViewUserProfile retrieves the user's profile information from the API using the provided token.
// UserProfile represents the structure of user profile data
type UserProfile struct {
	ID       string `json:"user_id"`
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}

// ViewUserProfile retrieves the user's profile information from the API using the provided token.
func ViewUserProfile(token string) (string, string, string, string, error) {
	apiURL := "http://localhost:8981/task_app/user_manager_service/api/v0.1/user/verify"
	client := &http.Client{}

	req, err := http.NewRequestWithContext(context.Background(), "GET", apiURL, nil)
	if err != nil {
		return "", "", "", "", fmt.Errorf("error creating request: %w", err)
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json") // Ensure the content type is set to JSON

	resp, err := client.Do(req)
	if err != nil {
		return "", "", "", "", fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", "", "", "", fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", "", "", fmt.Errorf("error reading response body: %w", err)
	}

	// Log the raw response body for debugging
	//fmt.Printf("Raw response body: %s\n", body)

	// Parse the JSON response
	var userProfile UserProfile
	if err := json.Unmarshal(body, &userProfile); err != nil {
		return "", "", "", "", fmt.Errorf("error parsing response body: %w", err)
	}

	return userProfile.ID, userProfile.UserName, userProfile.Name, userProfile.Email, nil
}
