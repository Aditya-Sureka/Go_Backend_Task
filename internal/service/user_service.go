package service

import (
	"context"
	"time"

	"github.com/Aditya-Sureka/Go_Backend_Task/db/sqlc"
	"github.com/Aditya-Sureka/Go_Backend_Task/internal/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(
	repo *repository.UserRepository,
) *UserService {
	return &UserService{
		Repo: repo,
	}
}

func (s *UserService) CreateUser(
	ctx context.Context,
	name string,
	dob string,
) (sqlc.User, error) {

	parsedDOB, err := time.Parse(
		"2006-01-02",
		dob,
	)

	if err != nil {
		return sqlc.User{}, err
	}

	return s.Repo.CreateUser(
		ctx,
		name,
		parsedDOB,
	)
}