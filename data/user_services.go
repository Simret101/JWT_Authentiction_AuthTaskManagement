package data

import (
	"errors"
	"fmt"
	"sync"
	"task/config"
	"task/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var (
	users      = []models.User{}
	lastUserID = 0
	userMu     sync.Mutex
)


func CreateUser(user *models.User) error {
	userMu.Lock()
	defer userMu.Unlock()

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
	lastUserID++
	user.ID = lastUserID

	users = append(users, *user)
	fmt.Printf("User created: %v\n", user)
	return nil
}


func AuthenticateUser(username, password string) (string, error) {
	userMu.Lock()
	defer userMu.Unlock()

	for _, user := range users {
		if user.Username == username {
			if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err == nil {
				return generateToken(user)
			}
			break
		}
	}
	return "", errors.New("invalid username or password")
}

e
func GetUserByUsername(username string) (*models.User, error) {
	userMu.Lock()
	defer userMu.Unlock()

	for _, user := range users {
		if user.Username == username {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

func generateToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.ID,
		"role":   user.Role,
		"exp":    time.Now().Add(config.TokenExpiration).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.SecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(tokenString string) (*models.Claims, error) {
	claims := &models.Claims{}

	
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

	if claims.ExpiresAt < time.Now().Unix() {
		return nil, errors.New("token has expired")
	}

	return claims, nil
}

