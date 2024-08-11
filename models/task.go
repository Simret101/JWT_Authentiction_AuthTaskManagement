package models

import (
	"errors"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskStatus string

const (
	TaskStatusComplete   TaskStatus = "complete"
	TaskStatusInProgress TaskStatus = "in_progress"
	TaskStatusStarted    TaskStatus = "started"
)

type Task struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	DueDate     time.Time          `json:"duedate" bson:"duedate"`
	Status      TaskStatus         `json:"status" bson:"status"`
	UserID      int                `json:"userID" bson:"userID"`
}

func (t *Task) Validate() error {
	if t.ID == primitive.NilObjectID {
		t.ID = primitive.NewObjectID()
	}
	if err := validateTitle(t.Title); err != nil {
		return err
	}
	if err := validateDescription(t.Description); err != nil {
		return err
	}
	if err := validateDueDate(t.DueDate); err != nil {
		return err
	}
	if err := validateStatus(t.Status); err != nil {
		return err
	}
	return nil
}

func validateTitle(title string) error {
	if strings.TrimSpace(title) == "" {
		return errors.New("title must not be empty")
	}
	if len(title) > 100 {
		return errors.New("title must be less than 100 characters")
	}
	if len(title) < 3 {
		return errors.New("title must be greater than 3 characters")
	}
	return nil
}

func validateDescription(description string) error {
	if strings.TrimSpace(description) == "" {
		return errors.New("description must not be empty")
	}
	return nil
}

func validateDueDate(dueDate time.Time) error {
	if dueDate.IsZero() {
		return errors.New("due date must be specified")
	}
	return nil
}

func validateStatus(status TaskStatus) error {
	switch status {
	case TaskStatusComplete, TaskStatusInProgress, TaskStatusStarted:
		return nil
	default:
		return errors.New("status is invalid")
	}
}
