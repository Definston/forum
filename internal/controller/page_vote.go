package controller

import (
	"net/http"

	"forum/internal/model"
)

func (s *server) vote(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/vote" {
		s.error(w, http.StatusNotFound, nil)
		return
	}

	uid := r.Context().Value(model.ContextKeySession).(int)

	if uid <= 0 {
		s.error(w, http.StatusUnauthorized, nil)
		return
	}

	switch r.Method {
	case http.MethodPost:
		pid := r.FormValue("post_id")
		vote := r.FormValue("vote")

		if err := s.service.Vote.Vote(uid, pid, vote); err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		http.Redirect(w, r, r.FormValue("path"), http.StatusFound)
		return

	default:
		s.error(w, http.StatusMethodNotAllowed, nil)
		return
	}
}
