package session

import (
	"net/http"
	"time"

	"github.com/gofrs/uuid"
)

func (s *session) AddSession(w http.ResponseWriter, uid int) error {
	token, err := uuid.DefaultGenerator.NewV4()
	if err != nil {
		return err
	}

	err = s.rsess.AddSession(token.String(), uid, time.Now().Add(timeLimitExpire))
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:  cookieName,
		Value: token.String(),
	})
	return nil
}
