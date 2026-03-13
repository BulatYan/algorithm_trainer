package models

type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	PasswordHash string `json:"password_hash"`
	Email        string `json:"email"`
}

type Task struct {
	ID          int      `json:"id"`
	ID_USER     int      `json:"id_user"`
	Name        string   `json:"name"`
	Input_data  []string `json:"input_data"`
	Output_data []string `json:"output_data"`
	Description string   `json:"description"`
	Lvl         int      `json:"lvl"`
}
