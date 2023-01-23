package session

import (
	"context"
	"net/http"
	"time"

	"forum/internal/model"
)

// return -1 if session is not exist or is expire, refrech expire time if exist
func (s *session) GetSession(r *http.Request) (context.Context, error) {
	ck, err := r.Cookie(cookieName)
	ctx := context.WithValue(context.Background(), model.ContextKeySession, -1)

	switch err {
	case nil:
		uid, expireTime, err := s.rsess.GetSession(ck.Value)
		if err != nil {
			return ctx, err
		}

		if uid == 0 {
			return ctx, nil
		}

		if expireTime.Before(time.Now()) {
			if err = s.rsess.DelSession(uid); err != nil {
				return ctx, err
			}
			return ctx, nil
		}

		if err = s.rsess.RefreshTime(uid, time.Now().Add(timeLimitExpire)); err != nil {
			return ctx, err
		}

		ctx := context.WithValue(context.Background(), model.ContextKeySession, uid)
		return ctx, nil

	case http.ErrNoCookie:
		return ctx, nil

	default:
		return ctx, err
	}
}
