package data

import (
	"errors"
	"fmt"
	"sync"
	"task/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	tasks = []models.Task{}
	mu    sync.Mutex
)


func GetAllTasks() []models.Task {
	mu.Lock()
	defer mu.Unlock()
	return tasks
}


func GetTasksByUserID(userID int) []models.Task {
	mu.Lock()
	defer mu.Unlock()
	var userTasks []models.Task
	for _, task := range tasks {
		if task.UserID == userID {
			userTasks = append(userTasks, task)
		}
	}
	return userTasks
}


func GetTaskByID(id primitive.ObjectID) (*models.Task, error) {
	mu.Lock()
	defer mu.Unlock()
	for _, task := range tasks {
		if task.ID == id {
			return &task, nil
		}
	}
	return nil, errors.New("task not found")
}


func CreateTask(task *models.Task) {
	mu.Lock()
	defer mu.Unlock()
	task.ID = primitive.NewObjectID() // Generate a new ObjectID
	tasks = append(tasks, *task)
	fmt.Printf("Task created: %v\n", task)
}

// UpdateTask updates an existing task by its ObjectID.
func UpdateTask(id primitive.ObjectID, updatedTask *models.Task) error {
	mu.Lock()
	defer mu.Unlock()
	for i, task := range tasks {
		if task.ID == id {
			updatedTask.ID = id
			tasks[i] = *updatedTask
			fmt.Printf("Task updated: %v\n", updatedTask)
			return nil
		}
	}
	return errors.New("task not found")
}


func DeleteTask(id primitive.ObjectID) error {
	mu.Lock()
	defer mu.Unlock()
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Printf("Task deleted with ID: %s\n", id.Hex())
			return nil
		}
	}
	return errors.New("task not found")
}

