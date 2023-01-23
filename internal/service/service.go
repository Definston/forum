package service

import (
	"forum/internal/model"
	"forum/internal/repository"
	"forum/internal/service/auth"
	"forum/internal/service/post"
	"forum/internal/service/reg"
	"forum/internal/service/session"
	"forum/internal/service/tag"
	"forum/internal/service/vote"
)

type Service struct {
	Auth model.AuthService
	Reg  model.RegService
	Sess model.SessServise
	Post model.PostService
	Vote model.VoteService
	Tag  model.TagService
}

func NewService(repo *repository.Repository) *Service {
	auth := auth.NewServiceAuth(repo.User)
	reg := reg.NewServiceReg(repo.User)
	tag := tag.NewServiceTag(repo.Tag)
	sess := session.NewServiceSession(repo.Session)
	vote := vote.NewServiceVote(repo.Vote)
	post := post.NewServicePost(repo.Post, tag, vote)

	return &Service{
		Auth: auth,
		Reg:  reg,
		Tag:  tag,
		Sess: sess,
		Vote: vote,
		Post: post,
	}
}
