package repository

import (
	"database/sql"

	"go.uber.org/zap"
)

type Repository struct {
	AssignmentRepo AssignmentRepository
	SubmissionRepo SubmissionRepo
	UserRepo       UserRepository
}

func NewRepository(db *sql.DB, log *zap.Logger) Repository {
	return Repository{
		AssignmentRepo: NewAssignmentRepository(db),
		SubmissionRepo: NewSubmissionRepo(db),
		UserRepo:       NewUserRepository(db, log),
	}
}
