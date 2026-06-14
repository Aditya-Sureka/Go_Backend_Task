package repository

import (
	"context"
	"time"

	"github.com/Aditya-Sureka/Go_Backend_Task/db/sqlc"
)

type UserRepository struct {
	Queries *sqlc.Queries
}

func (r *UserRepository) GetUser(
	ctx context.Context,
	id int32,
) (sqlc.User, error) {
	return r.Queries.GetUser(
		ctx,
		id,
	)
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

func (r *UserRepository) ListUsers(
ctx context.Context,
) ([]sqlc.User, error) {


return r.Queries.ListUsers(ctx)
}

func (r *UserRepository) UpdateUser(
ctx context.Context,
id int32,
name string,
dob time.Time,
) (sqlc.User, error) {


return r.Queries.UpdateUser(
	ctx,
	sqlc.UpdateUserParams{
		ID:   id,
		Name: name,
		Dob:  dob,
	},
)
}

func (r *UserRepository) DeleteUser(
ctx context.Context,
id int32,
) error {


return r.Queries.DeleteUser(
	ctx,
	id,
)


}


