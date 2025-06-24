package model

type User struct {
	Model
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Role     string `json:"rule" validate:"required,oneof=student lecture admin"`
}
