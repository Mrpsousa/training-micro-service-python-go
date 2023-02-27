package model
 
type User struct {
    Id   string `json:"id"`
    Name string `json:"user_name"`
    City string `json:"city"`
	Cpf int64    `json:"cpf"`
	Email string `json:"email"`
	Phone string    `json:"phone"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []User
}