package config

import "time"

// defines the config variables
var (
	SecretKey       = "my_secret_key"
	TokenExpiration = 1 * time.Hour
	DatabaseURI     = "mongodb://localhost:27017"
	UserDBName      = "userdb"
	TaskDBName      = "taskdb"
)
