package repository

import (
	"context"
	"database/sql"
	"fmt"
	"test-crud-api/internal/model"
	"time"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string, store *Storage) (*Storage, error) {

	db, err := sql.Open("postgres", storagePath)
	if err != nil {
		return nil, fmt.Errorf("нет ссоединения с базой", err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("нет ссоединения с базаой", err)
	}
	return &Storage{db: db}, nil
}

func (s *Storage) GetUserById(ctx context.Context, id string) (model.User, error) {
	u := model.User{}
	err := s.db.QueryRow("SELECT id,firstname,lastname,age,recordingdate FROM Users WHERE id = $1", sql.Named("id", id)).Scan(&u.ID, &u.FirstName, &u.LastName, &u.Age, &u.RecordingDate)
	if err != nil {
		return u, err
	}
	return u, nil
}

func (s *Storage) CreateUser(ctx context.Context, u model.User, t time.Time) error {

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	createUserQwery := `INSERT INTO  Users (id,  first_name, last_name, age, recording_date) VALUES ($1, $2, $3 ,$4 ,$5) RETURNING id`
	row := tx.QueryRow(createUserQwery, u.ID, u.FirstName, u.LastName, u.Age, t)
	fmt.Println(row.Err())
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

		err = rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Age, &u.RecordingDate)
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

/*
	func (s* Storage)GetAllUsersWithFilters(filterOptions []Field,)(,error){
	   if []fields

}
*/
func (s *Storage) DeleteUser(ctx context.Context, id string) error {

	_, err := s.db.Exec("DELETE FROM User WHERE  id = :id", sql.Named("number", id))
	if err != nil {
		return err
	}
	return nil
}
