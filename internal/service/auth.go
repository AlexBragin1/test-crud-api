package service

import (
	"context"
	"test-crud-api/internal/lib"
	"test-crud-api/internal/model"
	"test-crud-api/internal/repository"
	"time"
)

type Service struct {
	store repository.Repo
}

func NewService(store repository.Repo) *Service {
	return &Service{store: store}
}

func (s *Service) CreateUser(ctx context.Context, user model.User) error {
	var t time.Time
	if user.ID == "" || len(user.ID) < 10 {
		user.ID = lib.NewRandomString(10)
	}
	if user.RecordingDate == 0 {
		t = time.Now()
	}
	t = time.Unix(0, user.RecordingDate)

	return s.store.CreateUser(ctx, user, t)
}

func (s *Service) GetUserByID(ctx context.Context, id string) (model.User, error) {
	return s.store.GetUserById(ctx, id)
}

//func (s *Service) GetAll(ctx *context.Context) ([]model.User, error) {
//	return s.store.getAllUsersWithFilters(ctx)
//}

func (s *Service) DeleteUser(ctx context.Context, id string) error {
	return s.store.DeleteUser(ctx, id)
}
