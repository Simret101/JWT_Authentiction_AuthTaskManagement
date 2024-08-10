package config

import "time"

var (
	TokenExpiration = 1 * time.Hour
	DatabaseURI     = "mongodb://localhost:27017"
	UserDBName      = "userdb"
	TaskDBName      = "taskdb"
)

const (
	ServerPort = ":8080"
	SecretKey  = "your_secret_key"
)
