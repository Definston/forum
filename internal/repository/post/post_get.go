package post

import (
	"database/sql"

	"forum/internal/model"
)

func (p *post) GetPostById(pid int) (*model.Post, error) {
	q := `SELECT * FROM posts WHERE id == ?`
	row := p.db.QueryRow(q, pid)

	switch row.Err() {
	case nil:
		post := model.Post{}
		if err := row.Scan(&post.Id, &post.UserId, &post.UserNick, &post.ParentId, &post.Content, &post.Link, &post.Resiever); err != nil {
			return nil, err
		}

		return &post, nil
	default:
		return nil, row.Err()
	}
}

func (p *post) GetPostsByParentId(pid int) (*[]model.Post, error) {
	q := `SELECT * FROM posts WHERE parent_id == ?`
	rows, err := p.db.Query(q, pid)

	return p.fillPost(rows, err)
}

func (p *post) GetPostsByTag(tag string) (*[]model.Post, error) {
	q := `SELECT * FROM posts WHERE parent_id == 0 AND id IN (SELECT post_id FROM tags WHERE tag == ?)`
	rows, err := p.db.Query(q, tag)

	return p.fillPost(rows, err)
}

func (p *post) GetPostsLikedByUser(uid int) (*[]model.Post, error) {
	q := `SELECT * FROM posts WHERE parent_id == 0 AND id IN (SELECT post_id FROM votes WHERE user_id == ? AND vote == 1)`
	rows, err := p.db.Query(q, uid)

	return p.fillPost(rows, err)
}

func (p *post) GetPostsAddedByUser(uid int) (*[]model.Post, error) {
	q := `SELECT * FROM posts WHERE user_id == ? AND parent_id == 0`
	rows, err := p.db.Query(q, uid)

	return p.fillPost(rows, err)
}

func (p *post) fillPost(rows *sql.Rows, err error) (*[]model.Post, error) {
	switch err {
	case nil:
		posts := []model.Post{}
		for rows.Next() {
			post := model.Post{}
			if err := rows.Scan(&post.Id, &post.UserId, &post.UserNick, &post.ParentId, &post.Content, &post.Link, &post.Resiever); err != nil {
				return nil, err
			}
			posts = append(posts, post)
		}
		return &posts, nil

	case sql.ErrNoRows:
		return &[]model.Post{}, nil
	default:
		return nil, err
	}
}
