package data

import (
	"errors"
	"fmt"
	"task/config"
	"task/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var users = []models.User{}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// adds a new user to the database
func CreateUser(user *models.User) error {
	for _, u := range users {
		if u.Username == user.Username {
			return errors.New("username already exists")
		}
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	users = append(users, *user)
	fmt.Printf("User created: %v\n", user)
	return nil
}

// verifies user credentials and returns a JWT token
func AuthenticateUser(username, password string) (string, error) {
	for _, user := range users {
		if user.Username == username {
			if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
				return "", errors.New("invalid credentials")
			}
			return generateJWT(username)
		}
	}
	return "", errors.New("invalid credentials")
}

// checks the validity of the JWT token
func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SecretKey), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("invalid token signature")
		}
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}

// Generatesa JWT
func generateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(config.TokenExpiration)
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.SecretKey))
}
