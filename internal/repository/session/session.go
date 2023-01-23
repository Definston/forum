package session

import (
	"net/http"
	"time"
)

type session struct {
	st map[int]*http.Cookie
}

func NewRepositorySession() *session {
	return &session{
		st: make(map[int]*http.Cookie),
	}
}

func (s *session) AddSession(token string, uid int, time time.Time) error {
	s.st[uid] = &http.Cookie{
		Value:   token,
		Expires: time,
	}
	return nil
}

func (s *session) GetSession(token string) (int, time.Time, error) {
	for uid, cookie := range s.st {
		if cookie.Value == token {
			return uid, cookie.Expires, nil
		}
	}

	return 0, time.Time{}, nil
}

func (s *session) DelSession(uid int) error {
	delete(s.st, uid)
	return nil
}

func (s *session) RefreshTime(uid int, time time.Time) error {
	s.st[uid].Expires = time
	return nil
}
