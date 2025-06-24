package service

import (
	"errors"
	"go-23/model"
	"go-23/repository"

	"go.uber.org/zap"
)

type AuthService interface {
	Login(email, password string) (*model.User, error)
}

type authService struct {
	Repo repository.Repository
	Log  *zap.Logger
}

func NewAuthService(repo repository.Repository, log *zap.Logger) AuthService {
	return &authService{
		Repo: repo,
		Log:  log,
	}
}

func (s *authService) Login(email, password string) (*model.User, error) {
	// fmt.Println(email, password)
	user, err := s.Repo.UserRepo.FindByEmail(email)
	if err != nil {
		s.Log.Error("Service auth", zap.String("error", err.Error()))
		return nil, errors.New("user not found")
	}

	if user.Password != password {
		return nil, errors.New("incorrect password")
	}

	return user, nil
}
