package models

type User struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	UserType string `json:"user_type"`
}
