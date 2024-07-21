package core

import (
	"crypto/rand"
	"encoding/base64"
	"log"

	//"go/token"
	"time"

	"github.com/dgrijalva/jwt-go"

	//"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var secret_key, _ = generateSecretKey(64)

var Jwtkey = []byte(secret_key)

type JwtClaim struct {
	UserID   string
	UserName string
	Email    string
	jwt.StandardClaims
}

func GenerateToken(user_id string, user_name string, email string) (string, error) {
	expirationTIme := time.Now().Add(24 * time.Hour)
	claims := &JwtClaim{
		UserID:   user_id,
		UserName: user_name,
		Email:    email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTIme.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(Jwtkey)
}
func ValidateToken(tokenString string) (*JwtClaim, error) {
	claims := &JwtClaim{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return Jwtkey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}
	return claims, nil
}

func generateSecretKey(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}
func HashPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil

}

func CheckPasswordHash(password, hash string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			log.Println("Password mismatch")
		} else {
			log.Println("Error checking password:", err)
		}
		return false
	}
	return true
}
