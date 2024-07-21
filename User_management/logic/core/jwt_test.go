package core

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	userID := "12345"
	userName := "testuser"
	email := "testuser@example.com"

	token, err := GenerateToken(userID, userName, email)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestValidateToken(t *testing.T) {
	userID := "12345"
	userName := "testuser"
	email := "testuser@example.com"

	token, err := GenerateToken(userID, userName, email)
	assert.NoError(t, err)

	claims, err := ValidateToken(token)
	assert.NoError(t, err)
	assert.Equal(t, userID, claims.UserID)
	assert.Equal(t, userName, claims.UserName)
	assert.Equal(t, email, claims.Email)
}

func TestExpiredToken(t *testing.T) {
	// Create a token that expires immediately
	expiredTime := time.Now().Add(-time.Hour)
	claims := &JwtClaim{
		UserID:   "test_user_id",
		UserName: "test_user_name",
		Email:    "test@example.com",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(Jwtkey)
	assert.NoError(t, err)

	// Validate the token
	_, err = ValidateToken(tokenString)
	if ve, ok := err.(*jwt.ValidationError); ok {
		assert.Equal(t, ve.Errors, jwt.ValidationErrorExpired)
	} else {
		t.Errorf("expected jwt.ValidationError, got %T", err)
	}
}

func TestHashPassword(t *testing.T) {
	password := "mysecretpassword"

	hash, err := HashPassword(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)

	// Check if the hashed password is valid
	isValid := CheckPasswordHash(password, hash)
	assert.True(t, isValid)
}

func TestCheckPasswordHash(t *testing.T) {
	password := "Passwrod123!"
	wrongPassword := "wrongpassword"

	hash, err := HashPassword(password)
	assert.NoError(t, err)

	// Check if the correct password matches the hash
	isValid := CheckPasswordHash(password, hash)
	assert.True(t, isValid)

	// Check if the wrong password does not match the hash
	isValid = CheckPasswordHash(wrongPassword, hash)
	assert.False(t, isValid)
}
