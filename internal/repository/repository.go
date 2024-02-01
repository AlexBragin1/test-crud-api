package repository

import (
	"context"
	"database/sql"
	"fmt"
	"test-crud-api/internal/model"
	"test-crud-api/pkg/filter"
	"time"
)

type Repo interface {
	GetUserById(ctx context.Context, id string) (model.User, error)
	CreateUser(ctx context.Context, user model.User, t time.Time) error
	GetAllUsersWithFilters(ctx context.Context, filterOptions filter.Options) ([]model.User, error)
	FindAllUsers(ctx context.Context) error
	DeleteUser(ctx context.Context, id string) error
}
type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewDB(cfg Config) *sql.DB {

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return db

}
