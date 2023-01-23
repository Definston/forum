package post

import "forum/internal/model"

func (p *post) AddPost(post *model.Post) (int, error) {
	q := `INSERT INTO posts (user_id, user_nick, content) VALUES (?, ?, ?)`
	row, err := p.db.Exec(q, post.UserId, post.UserNick, post.Content)
	if err != nil {
		return 0, err
	}

	id, err := row.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (p *post) AddComm(post *model.Post) (int, error) {
	q := `INSERT INTO posts (user_id, user_nick, content, parent_id) VALUES (?, ?, ?, ?)`
	row, err := p.db.Exec(q, post.UserId, post.UserNick, post.Content, post.ParentId)
	if err != nil {
		return 0, err
	}

	id, err := row.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (p *post) AddLinkById(pid int, link string) error {
	q := `UPDATE posts SET link = ? WHERE id = ?`
	_, err := p.db.Exec(q, link, pid)
	if err != nil {
		return err
	}

	return nil
}

func (p *post) AddResieverById(pid, resiever int) error {
	q := `UPDATE posts SET resiever = ? WHERE id = ?`
	_, err := p.db.Exec(q, resiever, pid)
	if err != nil {
		return err
	}

	return nil
}
