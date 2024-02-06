package repository

import (
	"context"
	"fmt"
	"test-crud-api/internal/model"
	"test-crud-api/pkg/filter"
	"time"

	"github.com/jmoiron/sqlx"
)

type Storage struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Storage {

	return &Storage{db: db}

}

func (s *Storage) GetUserById(ctx context.Context, id string) (model.User, error) {
	u := model.User{}
	tx, err := s.db.Begin()
	if err != nil {
		fmt.Println("Error 1")
		return u, err
	}
	var t time.Time
	query := "select id,firstname,lastname,age,recordingdate from users where id=$1"
	errQwery := tx.QueryRowContext(ctx, query, u.ID).Scan(&u.ID, &u.FirstName, &u.LastName, &u.Age, &t)
	if errQwery != nil {
		tx.Rollback()
		return u, err
	}

	u.RecordingDate = t.Unix()
	return u, nil
}

func (s *Storage) CreateUser(ctx context.Context, u model.User, t time.Time) error {
	fmt.Println(u)
	tx, err := s.db.Begin()
	if err != nil {
		fmt.Println("Error 1")
		return err
	}
	createUserQwery := "INSERT INTO  users (id, first_name, last_name, age, recording_date) VALUES ($1, $2, $3 ,$4 ,$5)"
	_, errExec := tx.ExecContext(ctx, createUserQwery, u.ID, u.FirstName, u.LastName, u.Age, t.Format("2001-09-29 00:00:00"))
	if errExec != nil {
		fmt.Println("Eror2")
		tx.Rollback()
		return err
	}

	return nil
}

func (s *Storage) FindAllUsers(ctx context.Context) ([]model.User, error) {

	tx, err := s.db.Begin()
	if err != nil {
		fmt.Println("Error 1")
		return nil, err
	}
	qwery := "select id,first_name,last_name,age,recording_date from users"
	rows, errQwery := tx.QueryContext(ctx, qwery)
	if errQwery != nil {
		tx.Rollback()
		return nil, err
	}
	defer rows.Close()
	var users []model.User
	for rows.Next() {
		var u model.User
		var t time.Time
		err = rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Age, &t)

		if err != nil {
			return nil, err
		}
		u.RecordingDate = t.Unix()
		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (s *Storage) GetAllUsersWithFilter(ctx context.Context, filterOptions filter.Options) ([]model.User, error) {

	tx, err := s.db.Begin()
	if err != nil {
		fmt.Println("Error 1")
		return nil, err
	}
	qwery := fmt.Sprintf("delete id,first_name,last_name,age,recording_date from users where %s", filterOptions.Fields())
	rows, errQwery := tx.QueryContext(ctx, qwery)
	if errQwery != nil {
		tx.Rollback()
		return nil, err
	}
	defer rows.Close()
	var users []model.User
	for rows.Next() {
		var u model.User
		var t time.Time
		err = rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Age, &t)

		if err != nil {
			return nil, err
		}
		u.RecordingDate = t.Unix()
		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (s *Storage) DeleteUser(ctx context.Context, id string) error {
	tx, err := s.db.Begin()
	if err != nil {
		fmt.Println("Error 1")
		return err
	}
	qwery := "delete from users where id=$1"
	_, errQwery := tx.ExecContext(ctx, qwery, id)
	if errQwery != nil {
		tx.Rollback()
		return err
	}

	return nil
}
