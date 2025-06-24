package service

import (
	"go-23/repository"

	"go.uber.org/zap"
)

type Service struct {
	AssignmentService AssignmentService
	SubmissionService SubmissionService
	UserService       UserService
	AuthService       AuthService
}

func NewService(repo repository.Repository, log *zap.Logger) Service {
	return Service{
		AssignmentService: NewAssignmentService(repo),
		SubmissionService: NewSubmissionService(repo),
		UserService:       NewUserService(repo),
		AuthService:       NewAuthService(repo, log),
	}
}
