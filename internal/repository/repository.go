package repository

import (
	"context"
	"fmt"
	"test-crud-api/internal/model"
	"test-crud-api/pkg/filter"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repo interface {
	GetUserById(ctx context.Context, id string) (model.User, error)
	CreateUser(ctx context.Context, user model.User, t time.Time) error
	GetAllUsersWithFilter(ctx context.Context, filterOptions filter.Options) ([]model.User, error)
	FindAllUsers(ctx context.Context) error
	DeleteUser(ctx context.Context, id string) error
}
type Config struct {
	Host     string
	Port     string
	Username string
	DBName   string
	SSLMode  string
	Password string
}

func NewDB(cfg Config) *sqlx.DB {

	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		fmt.Println("don't open db", err)
		return nil
	}
	defer db.Close()
	fmt.Println("open db")
	err = db.Ping()
	if err != nil {
		fmt.Println("don't ping db", err)
		return nil
	}
	return db

}
