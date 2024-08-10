package models


type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"duedate"`
	Status      string `json:"status"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
	UserID int    `json:"userID"`
}
