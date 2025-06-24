package dto

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type GradeSubmitRequest struct {
	UserID       int     `json:"user_id" validate:"required"`
	AssignmentID int     `json:"assignment_id" validate:"required"`
	Grade        float64 `json:"grade" validate:"required"`
}
