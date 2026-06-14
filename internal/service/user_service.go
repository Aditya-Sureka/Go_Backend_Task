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

func (s *UserService) GetUser(
	ctx context.Context,
	id int32,
) (sqlc.User, error) {
	return s.Repo.GetUser(
		ctx,
		id,
	)
}

func (s *UserService) ListUsers(
ctx context.Context,
) ([]sqlc.User, error) {

return s.Repo.ListUsers(ctx)
}

func (s *UserService) UpdateUser(
ctx context.Context,
id int32,
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

return s.Repo.UpdateUser(
	ctx,
	id,
	name,
	parsedDOB,
)
}
