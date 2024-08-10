package controllers

import (
	"net/http"
	"strconv"
	"task/data"
	"task/models"

	"github.com/gin-gonic/gin"
)

func GetAllTasks(c *gin.Context) {
	role := c.GetString("role")
	userID := c.GetInt("userID")

	var tasks []models.Task

	if role == "admin" {
		tasks = data.GetAllTasks()
	} else {
		tasks = data.GetTasksByUserID(userID)
	}

	c.JSON(http.StatusOK, tasks)
}

func GetTaskByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	role := c.GetString("role")
	userID := c.GetInt("userID")

	task, err := data.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	if role != "admin" && task.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to view this task"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	task.UserID = c.GetInt("userID")
	data.CreateTask(&task)
	c.JSON(http.StatusCreated, task)
}

func UpdateTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	role := c.GetString("role")
	userID := c.GetInt("userID")

	task, err := data.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	if role != "admin" && task.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to update this task"})
		return
	}

	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	updatedTask.ID = id
	updatedTask.UserID = task.UserID

	if err := data.UpdateTask(id, &updatedTask); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	c.JSON(http.StatusOK, updatedTask)
}

func DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	role := c.GetString("role")
	userID := c.GetInt("userID")

	task, err := data.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	if role != "admin" && task.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to delete this task"})
		return
	}

	if err := data.DeleteTask(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

