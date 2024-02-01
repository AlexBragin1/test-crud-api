package service

import (
	"context"
	"test-crud-api/internal/lib"
	"test-crud-api/internal/model"
	"test-crud-api/internal/repository"
	"test-crud-api/pkg/filter"
	"time"
)

type Service struct {
	store *repository.Storage
}

func NewService(store *repository.Storage) *Service {
	return &Service{store: store}
}

func (s *Service) CreateUser(ctx context.Context, user model.User) error {
	var t int64
	if user.ID == "" || len(user.ID) < 10 {
		user.ID = lib.NewRandomString(10)
	}
	if user.RecordingDate == 0 {
		t = time.Now().Unix()
	}
	return s.store.CreateUser(ctx, user, t)
}

func (s *Service) GetUserByID(ctx context.Context, id string) (model.User, error) {
	return s.store.GetUserById(ctx, id)
}
func (s *Service) FindAllUsers(ctx context.Context) ([]model.User, error) {
	return s.store.FindAllUsers(ctx)
}
func (s *Service) GetAllUsersWithFilters(ctx context.Context, filterOptions filter.Options) ([]model.User, error) {

	return s.store.GetAllUsersWithFilters(ctx, filterOptions)
}

func (s *Service) DeleteUser(ctx context.Context, id string) error {
	return s.store.DeleteUser(ctx, id)
}
