package data

import (
	"errors"

	"sync"
	"task/config"
	"task/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var (
	lastUserID = 0
	users      = []models.User{}
	userMu     sync.Mutex
)

func GenerateNewUserID() int {
	userMu.Lock()
	defer userMu.Unlock()
	lastUserID++
	return lastUserID
}

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

	users = append(users, *user)
	return nil
}

func AuthenticateUser(username, password string) (string, error) {
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
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.SecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
