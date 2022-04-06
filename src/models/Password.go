package models

type Password struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}
