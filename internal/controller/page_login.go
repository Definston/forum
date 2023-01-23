package controller

import (
	"database/sql"
	"errors"
	"html/template"
	"net/http"

	"forum/internal/model"
)

func (s *server) login(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
		s.error(w, http.StatusNotFound, nil)
		return
	}

	uid := r.Context().Value(model.ContextKeySession).(int)
	if uid > 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tmp, err := template.ParseFiles("./web/static/html/login.html")
	if err != nil {
		s.error(w, http.StatusInternalServerError, err)
		return
	}

	var message string

	switch r.Method {
	case http.MethodPost:

		user, err := s.service.Auth.AuthUser(r.FormValue("email"), r.FormValue("password"))
		switch err {
		case nil:
			if err := s.service.Sess.AddSession(w, user.Id); err != nil {
				s.error(w, http.StatusInternalServerError, err)
				return
			}

			http.Redirect(w, r, "/", http.StatusFound)

		case sql.ErrNoRows:
			message = "User is not exist!"

		case errors.New("wrong pass"):
			message = "Incorrect password!"

		default:
			s.error(w, http.StatusInternalServerError, err)
			return
		}

	case http.MethodGet:
	default:
		s.error(w, http.StatusMethodNotAllowed, nil)
		return
	}

	if err := tmp.ExecuteTemplate(w, "login", message); err != nil {
		s.error(w, http.StatusInternalServerError, err)
		return
	}
}
