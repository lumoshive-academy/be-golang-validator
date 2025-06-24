package service

import (
	"go-23/repository"
	"math"

	"go.uber.org/zap"
)

func totalPage(limit int, totalData int64) int {
	if totalData <= 0 {
		return 0
	}

	flimit := float64(limit)
	fdata := float64(totalData)

	res := math.Ceil(fdata / flimit)

	return int(res)
}

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
