package model

import (
	"context"
	"net/http"
)

type Service interface {
	AuthService
	RegService
	SessServise
	PostService
	VoteService
	TagService
}

type AuthService interface {
	AuthUser(email, pass string) (*User, error)
	GetUserById(uid int) (*User, error)

	IsValidUser(email, nick, pass string) bool
}

type RegService interface {
	AddUser(*User) (bool, error)
}

type SessServise interface {
	AddSession(w http.ResponseWriter, uid int) error
	GetSession(r *http.Request) (context.Context, error)
	DelSession(r *http.Request) error
}

type PostService interface {
	AddPost(*Post) error
	AddComm(*Post) error

	GetAllPosts() (*[]Post, error)
	GetPostById(pid int) (*Post, error)
	GetCommById(pid int) (*[]Post, error)
	GetPostsByTag(tag string) (*[]Post, error)
	GetPostsLikedByUser(uid int) (*[]Post, error)
	GetPostsAddedByUser(uid int) (*[]Post, error)

	IsValidContent(content string) bool
}

type VoteService interface {
	Vote(uid int, pid, vote string) error
	GetVoteByPostId(pid int) (likes, dislikes int, err error)
}

type TagService interface {
	GetTagAll() (*[]string, error)

	FindTags(string) (*map[string]bool, error)
	AddTagByPostId(pid int, tags *map[string]bool) error
	// ReplaceTagsToLink(content string, tags *map[string]bool) (string, error)
}
