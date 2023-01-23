package user

import (
	"database/sql"

	"forum/internal/model"
)

type user struct {
	db *sql.DB
}

func NewRepositoryUser(db *sql.DB) *user {
	return &user{
		db: db,
	}
}

func (u *user) AddUser(user *model.User) error {
	q := `INSERT INTO users (nick, email, pass) VALUES (?, ?, ?)`

	if _, err := u.db.Exec(q, user.Nick, user.Email, user.Pass); err != nil {
		return err
	}

	return nil
}

func (u *user) GetUserById(uid int) (*model.User, error) {
	q := `SELECT * FROM users WHERE id == ?`
	row := u.db.QueryRow(q, uid)

	switch row.Err() {
	case nil:
		user := &model.User{}
		err := row.Scan(&user.Id, &user.Email, &user.Nick, &user.Pass)
		if err != nil {
			return nil, err
		}
		return user, nil

	case sql.ErrNoRows:
		return nil, nil

	default:
		return nil, row.Err()
	}
}

func (u *user) GetUserByEmail(email string) (*model.User, error) {
	q := `SELECT * FROM users WHERE email == ?`
	row := u.db.QueryRow(q, email)

	switch row.Err() {
	case nil:
		user := &model.User{}
		err := row.Scan(&user.Id, &user.Email, &user.Nick, &user.Pass)
		if err != nil {
			return nil, err
		}
		return user, nil

	case sql.ErrNoRows:
		return nil, nil

	default:
		return nil, row.Err()
	}
}
