package repository

import (
	"context"
	"time"

	"github.com/student/go-user-api/db/sqlc"
	"github.com/student/go-user-api/internal/logger"
	"go.uber.org/zap"
)

type UserRepository struct {
	queries *sqlc.Queries
}

func NewUserRepository(queries *sqlc.Queries) *UserRepository {
	return &UserRepository{queries: queries}
}

func (r *UserRepository) Create(ctx context.Context, name string, dob time.Time) (sqlc.User, error) {
	logger.Log.Info("DB: Creating user", zap.String("name", name))
	return r.queries.CreateUser(ctx, sqlc.CreateUserParams{
		Name: name,
		Dob:  dob,
	})
}

func (r *UserRepository) GetByID(ctx context.Context, id int32) (sqlc.User, error) {
	logger.Log.Info("DB: Getting user by ID", zap.Int32("id", id))
	return r.queries.GetUserByID(ctx, id)
}

func (r *UserRepository) Update(ctx context.Context, id int32, name string, dob time.Time) (sqlc.User, error) {
	logger.Log.Info("DB: Updating user", zap.Int32("id", id))
	return r.queries.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:   id,
		Name: name,
		Dob:  dob,
	})
}

func (r *UserRepository) Delete(ctx context.Context, id int32) error {
	logger.Log.Info("DB: Deleting user", zap.Int32("id", id))
	return r.queries.DeleteUser(ctx, id)
}

func (r *UserRepository) List(ctx context.Context) ([]sqlc.User, error) {
	logger.Log.Info("DB: Listing all users")
	return r.queries.ListUsers(ctx)
}
