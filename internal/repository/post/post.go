package post

import "database/sql"

type post struct {
	db *sql.DB
}

func NewRepositoryPost(db *sql.DB) *post {
	return &post{
		db: db,
	}
}
