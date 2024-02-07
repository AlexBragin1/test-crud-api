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
	Db *sqlx.DB
}

func New(Db *sqlx.DB) Repo {

	return &Storage{Db: Db}

}
func (s *Storage) GetUserById(ctx context.Context, id string) (model.User, error) {
	u := model.User{}

	tx := s.Db.MustBegin()
	defer CommitOrRollback(tx)

	var t time.Time
	query := `SELECT 
	id,first_name,last_name,age
	FROM users 
	WHERE id=$1`
	errQwery := tx.QueryRowContext(ctx, query, u.ID).Scan(&u.ID, &u.FirstName, &u.LastName, &u.Age, &t)
	if errQwery != nil {

		return u, errQwery
	}

	u.RecordingDate = t.Unix()
	return u, nil
}
func (s *Storage) CreateUser(ctx context.Context, u model.User, t time.Time) error {
	tx := s.Db.MustBegin()
	defer CommitOrRollback(tx)
	createUserQwery := `INSERT INTO users (
		id, first_name, last_name, age, recording_date
		) VALUES (
			$1, $2, $3 ,$4 ,$5 
			)`
	tx.MustExecContext(ctx, createUserQwery, u.ID, u.FirstName, u.LastName, u.Age, t)
	return nil
}

func (s *Storage) GetAllUsersWithFilter(ctx context.Context, field filter.Field) ([]model.User, error) {
	var qwery string
	fmt.Println(field)
	tx := s.Db.MustBegin()
	defer CommitOrRollback(tx)
	if field.Name == "recording_dateTo" {
		qwery = fmt.Sprintf(`SELECT id, first_name, last_name, age, recording_date FROM users WHERE recording_date < to_timestamp(%s , 'DD/MM/YYYY') 
		and recording_date > to_timestamp(%s , 'DD/MM/YYYY')`, field.Operator, field.Value)
	} else if field.Operator == "=" {
		qwery = `SELECT 
	id, first_name, last_name, age, recording_date 
	FROM users`
	} else if field.Operator == "between" {
		qwery = fmt.Sprintf("SELECT id, first_name, last_name, age, recording_date FROM users WHERE %s %s %s", field.Name, myMap[field.Operator], field.Value)
	} else {

		qwery = fmt.Sprintf("SELECT id, first_name, last_name, age, recording_date FROM users WHERE %s%s%s", field.Name, myMap[field.Operator], field.Value)
	}
	fmt.Println(qwery)
	rows, errQwery := tx.QueryContext(ctx, qwery)
	if errQwery != nil {
		return nil, errQwery
	}

	defer rows.Close()
	var users []model.User
	for rows.Next() {
		var u model.User
		var t time.Time
		err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Age, &t)

		if err != nil {
			return nil, err
		}
		u.RecordingDate = t.Unix()
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil

}

func (s *Storage) FindAllUsers(ctx context.Context) ([]model.User, error) {
	tx := s.Db.MustBegin()

	defer CommitOrRollback(tx)
	qwery := `SELECT 
	id, first_name, last_name, age, recording_date 
	FROM users`
	rows, errQwery := tx.QueryContext(ctx, qwery)
	if errQwery != nil {
		fmt.Printf("don`t find allusers  %s", errQwery)
		return nil, errQwery
	}
	defer rows.Close()
	var users []model.User
	for rows.Next() {
		var u model.User
		var t time.Time
		err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Age, &t)

		if err != nil {
			return nil, err
		}
		u.RecordingDate = t.Unix()
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil

}

func (s *Storage) DeleteUser(ctx context.Context, id string) error {
	tx := s.Db.MustBegin()
	defer CommitOrRollback(tx)
	qwery := "delete from users where id=$1;"
	_, errQwery := tx.ExecContext(ctx, qwery, id)
	if errQwery != nil {
		fmt.Printf("don`t find id %s", errQwery)
		return errQwery
	}
	return nil
}
