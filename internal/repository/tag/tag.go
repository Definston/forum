package tag

import (
	"database/sql"
)

type tag struct {
	db *sql.DB
}

func NewRepositoryTag(db *sql.DB) *tag {
	return &tag{
		db: db,
	}
}

func (t *tag) AddTagByPostId(pid int, tag string) error {
	q := `INSERT INTO tags (post_id, tag) VALUES (?, ?)`
	_, err := t.db.Exec(q, pid, tag)
	if err != nil {
		return err
	}

	return nil
}

func (t *tag) GetTagAll() (*[]string, error) {
	q := `SELECT DISTINCT tag FROM tags`
	rows, err := t.db.Query(q)

	switch err {
	case nil:
		tags := []string{}
		for rows.Next() {
			var tag string
			if err := rows.Scan(&tag); err != nil {
				return nil, err
			}

			tags = append(tags, tag)
		}

		return &tags, nil

	case sql.ErrNoRows:
		return &[]string{}, nil
	default:
		return nil, err
	}
}
