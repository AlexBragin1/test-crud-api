package repository

import (
	"context"
	"database/sql"
	"fmt"
	"test-crud-api/internal/model"
	"test-crud-api/pkg/filter"
	"time"
)

type Storage struct {
	db *sql.DB
}

func New(db *sql.DB) *Storage {

	return &Storage{db: db}

}

func (s *Storage) GetUserById(ctx context.Context, id string) (model.User, error) {
	u := model.User{}
	err := s.db.QueryRow("SELECT id,firstname,lastname,age,recordingdate FROM Users WHERE id = $1", sql.Named("id", id)).Scan(&u.ID, &u.FirstName, &u.LastName, &u.Age, &u.RecordingDate)
	if err != nil {
		return u, err
	}
	return u, nil
}

func (s *Storage) CreateUser(ctx context.Context, u model.User, t int64) error {

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	createUserQwery := `INSERT INTO  Users (id,  first_name, last_name, age, recording_date) VALUES ($1, $2, $3 ,$4 ,$5) RETURNING id`
	row := tx.QueryRow(createUserQwery, u.ID, u.FirstName, u.LastName, u.Age, t)
	fmt.Println(row)
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
func (s *Storage) FindAllUsers(ctx context.Context) ([]model.User, error) {
	q := `SELECT id,  first_name, last_name, age, recording_date FROM Users`
	rows, err := s.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []model.User
	for rows.Next() {
		var u model.User
		var t time.Time
		err = rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Age, t)
		u.RecordingDate = t.Unix()
		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (s *Storage) GetAllUsersWithFilters(ctx context.Context, filterOptions []filter.Field) ([]model.User, error) {
	//	filter
	var u []model.User
	return u, nil
}

func (s *Storage) DeleteUser(ctx context.Context, id string) error {

	_, err := s.db.Exec("DELETE FROM User WHERE  id = :id", sql.Named("number", id))
	if err != nil {
		return err
	}
	return nil
}
