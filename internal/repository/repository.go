package repository

import (
	"database/sql"

	"forum/internal/model"
	"forum/internal/repository/post"
	"forum/internal/repository/session"
	"forum/internal/repository/tag"
	"forum/internal/repository/user"
	"forum/internal/repository/vote"
)

type Repository struct {
	Post model.PostRepo
	User model.UserRepo
	Vote model.VoteRepo
	// Relation model.RelationRepo
	Tag     model.TagRepo
	Session model.SessionRepo
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Post: post.NewRepositoryPost(db),
		User: user.NewRepositoryUser(db),
		Vote: vote.NewRepositoryVote(db),
		Tag:  tag.NewRepositoryTag(db),
		// Relation: relation.NewRepositoryRelation(db),
		Session: session.NewRepositorySession(),
	}
}
