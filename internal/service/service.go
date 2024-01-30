package service

import (
	"context"
	"test-crud-api/internal/model"
	"time"
)

type Servise interface {
	GetUserById(ctx context.Context, id string) (model.User, error)
	CreateUser(ctx context.Context, user model.User, t time.Time) error
	//GetAllUsersWithFilters(filterOptions []Field)([]User,error)
	FindAllUsers(ctx *context.Context) error
	DeleteUser(ctx *context.Context, id string) error
}
