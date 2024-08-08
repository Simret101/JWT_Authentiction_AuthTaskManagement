package models

//defines the user model
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

//defines the credentials model
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
