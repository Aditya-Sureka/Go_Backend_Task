package repository

import (
	"context"
	"time"

	"github.com/Aditya-Sureka/Go_Backend_Task/db/sqlc"
)

type UserRepository struct {
	Queries *sqlc.Queries
}

func NewUserRepository(q *sqlc.Queries) *UserRepository {
	return &UserRepository{
		Queries: q,
	}
}

func (r *UserRepository) CreateUser(
	ctx context.Context,
	name string,
	dob time.Time,
) (sqlc.User, error) {

	return r.Queries.CreateUser(
		ctx,
		sqlc.CreateUserParams{
			Name: name,
			Dob:  dob,
		},
	)
}