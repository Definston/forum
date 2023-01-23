package controller

import (
	"html/template"
	"net/http"

	"forum/internal/model"
)

func (s *server) signup(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/signup" {
		s.error(w, http.StatusNotFound, nil)
		return
	}

	uid := r.Context().Value(model.ContextKeySession).(int)
	if uid > 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tmp, err := template.ParseFiles("./web/static/html/signup.html")
	if err != nil {
		s.error(w, http.StatusInternalServerError, err)
		return
	}

	switch r.Method {
	case http.MethodPost:

		if !s.service.Auth.IsValidUser(r.FormValue("email"), r.FormValue("nick"), r.FormValue("password")) {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		user := model.User{
			Email: r.FormValue("email"),
			Nick:  r.FormValue("nick"),
			Pass:  r.FormValue("password"),
		}

		status, err := s.service.Reg.AddUser(&user)
		if err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		if status {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		message := "User already exist!"

		if err := tmp.ExecuteTemplate(w, "signup", message); err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

	case http.MethodGet:

		if err := tmp.ExecuteTemplate(w, "signup", nil); err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

	default:
		s.error(w, http.StatusMethodNotAllowed, nil)
		return
	}
}
