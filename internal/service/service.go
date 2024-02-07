package service

import (
	"context"
	"test-crud-api/internal/model"
	"test-crud-api/pkg/filter"
)

type Service interface {
	GetUserById(ctx context.Context, id string) (model.User, error)
	CreateUser(ctx context.Context, user model.User) error
	FindAllUsers(ctx context.Context) ([]model.User, error)
	GetAllUsersWithFilter(ctx context.Context, field filter.Field) ([]model.User, error)
	DeleteUser(ctx context.Context, id string) error
}
