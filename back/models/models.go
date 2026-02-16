package models

type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	PasswordHash string `json:"password_hash"`
	Email        string `json:"email"`
}
