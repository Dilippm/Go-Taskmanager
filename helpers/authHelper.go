package helpers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// Secret key used for signing JWT tokens
var jwtSecretKey = []byte("your_secret_key") // Replace with your actual secret key
func ComparePasswords(hashedPassword, password string) error {

	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// HashPassword hashes a plain-text password
func HashPassword(password string) (string, error) {
	// bcrypt generates a hashed version of the password with a cost of 14
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func GenerateJWTToken(id string, email string) (string, error) {
	// Define the expiration time of the token (e.g., 24 hours)
	expirationTime := time.Now().Add(24 * time.Hour)
	// Create a new token object
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   id,
		"email": email,                 // Subject (user ID)
		"exp":   expirationTime.Unix(), // Expiration time
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
