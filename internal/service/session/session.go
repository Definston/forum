package session

import (
	"time"

	"forum/internal/model"
)

const (
	cookieName      = "forum_session"
	timeLimitExpire = 300 * time.Second
)

type session struct {
	rsess model.SessionRepo
}

func NewServiceSession(r model.SessionRepo) *session {
	return &session{
		rsess: r,
	}
}
