package model

import "time"

type Repository interface {
	PostRepo
	UserRepo
	VoteRepo
	RelationRepo
	TagRepo
	SessionRepo
}

type PostRepo interface {
	AddPost(*Post) (pid int, err error)
	AddComm(*Post) (pid int, err error)
	AddLinkById(pid int, link string) error
	AddResieverById(pid, resiever int) error
	GetPostById(pid int) (*Post, error)
	GetPostsByParentId(pid int) (*[]Post, error)
	GetPostsByTag(tag string) (*[]Post, error)
	GetPostsLikedByUser(uid int) (*[]Post, error)
	GetPostsAddedByUser(uid int) (*[]Post, error)
}

type UserRepo interface {
	AddUser(*User) error
	GetUserById(uid int) (*User, error)
	GetUserByEmail(email string) (*User, error)
}

type VoteRepo interface {
	AddVote(pid, uid, vote int) error
	DelVote(pid, uid int) error
	GetVote(pid, uid int) (vote bool, err error)
	GetVoteByPostId(pid int) (likes, dislikes int64, err error)
}

type RelationRepo interface {
	AddRelation(answer int, resiever int) error
}

type TagRepo interface {
	AddTagByPostId(pid int, tag string) error
	GetTagAll() (*[]string, error)
}

type SessionRepo interface {
	AddSession(token string, uid int, time time.Time) error
	GetSession(token string) (int, time.Time, error)
	DelSession(uid int) error
	RefreshTime(uid int, time time.Time) error
}
