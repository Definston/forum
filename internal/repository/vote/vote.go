package vote

import "database/sql"

type vote struct {
	db *sql.DB
}

func NewRepositoryVote(db *sql.DB) *vote {
	return &vote{
		db: db,
	}
}

func (v *vote) AddVote(pid, uid, vote int) error {
	q := `INSERT INTO votes (post_id, user_id, vote) VALUES (?, ?, ?)`
	if _, err := v.db.Exec(q, pid, uid, vote); err != nil {
		return err
	}

	return nil
}

func (v *vote) DelVote(pid, uid int) error {
	q := `DELETE FROM votes WHERE post_id == ? AND user_id == ?`
	if _, err := v.db.Exec(q, pid, uid); err != nil {
		return err
	}

	return nil
}

func (v *vote) GetVote(pid, uid int) (vote bool, err error) {
	q := `SELECT vote FROM votes WHERE post_id == ? AND user_id == ?`
	row := v.db.QueryRow(q, pid, uid)

	if err = row.Scan(&vote); err != nil {
		return false, err
	}

	return vote, err
}

func (v *vote) GetVoteByPostId(pid int) (likes, dislikes int64, err error) {
	q := `SELECT count() FROM votes WHERE post_id == ? AND vote == 1`
	l := v.db.QueryRow(q, pid)
	q = `SELECT count() FROM votes WHERE post_id == ? AND vote == 0`
	d := v.db.QueryRow(q, pid)

	if err := l.Scan(&likes); err != nil {
		return 0, 0, err
	}
	if err := d.Scan(&dislikes); err != nil {
		return 0, 0, err
	}

	return likes, dislikes, nil
}
