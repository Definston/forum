package session

import "net/http"

func (s *session) DelSession(r *http.Request) error {
	ck, err := r.Cookie(cookieName)

	switch err {
	case nil:
		uid, _, err := s.rsess.GetSession(ck.Value)
		if err != nil {
			return err
		}

		if uid == 0 {
			return nil
		}

		if err = s.rsess.DelSession(uid); err != nil {
			return err
		}
		return nil

	case http.ErrNoCookie:
		return nil
	default:
		return err
	}
}
