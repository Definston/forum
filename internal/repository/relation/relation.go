package relation

import "database/sql"

type relation struct {
	db *sql.DB
}

func NewRepositoryRelation(db *sql.DB) *relation {
	return &relation{
		db: db,
	}
}

func (r *relation) AddRelation(pid int, resiever int) error {
	q := `INSERT INTO relations (resiever, answer) VALUES (?, ?)`
	if _, err := r.db.Exec(q, resiever, pid); err != nil {
		return err
	}

	return nil
}
