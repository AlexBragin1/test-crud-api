package service

import (
	"context"
	"fmt"
	"test-crud-api/internal/lib"
	"test-crud-api/internal/model"
	"test-crud-api/internal/repository"
	"test-crud-api/pkg/filter"
	"time"
)

type UserService struct {
	Repo repository.Repo
}

func NewService(store repository.Repo) Service {
	return &UserService{Repo: store}
}

func (s *UserService) GetUserById(ctx context.Context, id string) (model.User, error) {
	return s.Repo.GetUserById(ctx, id)
}

func (s *UserService) CreateUser(ctx context.Context, user model.User) error {
	var t time.Time
	if user.ID == "" || len(user.ID) != 10 {
		user.ID = lib.NewRandomString(10)
	}
	fmt.Println(user.ID)
	if user.RecordingDate == 0 {
		t = time.Now()
	} else {
		t = time.Unix(0, user.RecordingDate)
	}
	return s.Repo.CreateUser(ctx, user, t)
}

func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	return s.Repo.DeleteUser(ctx, id)
}

func (s *UserService) GetUserByID(ctx context.Context, id string) (model.User, error) {
	return s.Repo.GetUserById(ctx, id)
}

func (s *UserService) FindAllUsers(ctx context.Context) ([]model.User, error) {
	return s.Repo.FindAllUsers(ctx)
}

func (s *UserService) GetAllUsersWithFilter(ctx context.Context, field filter.Field) ([]model.User, error) {

	return s.Repo.GetAllUsersWithFilter(ctx, field)
}
