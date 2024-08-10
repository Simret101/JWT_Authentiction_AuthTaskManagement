package controllers

import (
	"net/http"
	"task/data"
	"task/models"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var creds models.Credentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	existingUser, _ := data.GetUserByUsername(creds.Username)
	if existingUser != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already taken"})
		return
	}

	role := "user"
	if creds.Role != "" {
		role = creds.Role
	}

	user := models.User{
		ID:       data.GenerateNewUserID(),
		Username: creds.Username,
		Password: creds.Password,
		Role:     role,
	}
	data.CreateUser(&user)
	c.JSON(http.StatusCreated, user)
}

func Login(c *gin.Context) {
	var credentials models.Credentials
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	token, err := data.AuthenticateUser(credentials.Username, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	user, _ := data.GetUserByUsername(credentials.Username)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"role":  user.Role,
	})
}
