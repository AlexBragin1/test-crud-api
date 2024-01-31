package service

import (
	"context"
	"test-crud-api/internal/model"
	"test-crud-api/pkg/filter"
	"time"
)

type Servise interface {
	GetUserById(ctx context.Context, id string) (model.User, error)
	CreateUser(ctx context.Context, user model.User, t time.Time) error
	FindAllUsers(ctx context.Context) ([]model.User, error)
	GetAllUsersWithFilters(ctx context.Context, filterOptions []filter.Field) ([]model.User, error)
	DeleteUser(ctx context.Context, id string) error
}
