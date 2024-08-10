package data

import (
	"errors"
	"fmt"
	"sync"
	"task/models"
)

var (
	tasks  = []models.Task{}
	lastID = 0
	mu     sync.Mutex
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

func GetTaskByID(id int) (*models.Task, error) {
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
	lastID++
	task.ID = lastID
	tasks = append(tasks, *task)
	fmt.Printf("Task created: %v\n", task)
}

func UpdateTask(id int, updatedTask *models.Task) error {
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

func DeleteTask(id int) error {
	mu.Lock()
	defer mu.Unlock()
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Printf("Task deleted with ID: %d\n", id)
			return nil
		}
	}
	return errors.New("task not found")
}

