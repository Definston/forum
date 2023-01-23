package controller

import (
	"net/http"

	"forum/internal/model"
)

func (s *server) logout(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/logout" {
		s.error(w, http.StatusNotFound, nil)
		return
	}

	uid := r.Context().Value(model.ContextKeySession).(int)
	if uid <= 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	switch r.Method {
	case http.MethodPost:

		if err := s.service.Sess.DelSession(r); err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		http.Redirect(w, r, "/", http.StatusFound)
		return

	default:
		s.error(w, http.StatusMethodNotAllowed, nil)
		return
	}
}
