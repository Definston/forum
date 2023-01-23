package controller

import (
	"html/template"
	"net/http"
	"strconv"

	"forum/internal/model"
)

func (s *server) post(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/post" {
		s.error(w, http.StatusNotFound, nil)
		return
	}

	uid := r.Context().Value(model.ContextKeySession).(int)

	pid, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		s.error(w, http.StatusBadRequest, err)
		return
	}

	data := struct {
		User  *model.User
		Post  *model.Post
		Comms *[]model.Post
	}{}

	user, err := s.service.Auth.GetUserById(uid)
	if err != nil {
		s.error(w, http.StatusInternalServerError, err)
		return
	}

	data.User = user

	switch r.Method {
	case http.MethodPost:
		if uid <= 0 {
			s.error(w, http.StatusUnauthorized, nil)
			return
		}

		if !s.service.Post.IsValidContent(r.FormValue("content")) {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		comm := model.Post{
			UserId:   user.Id,
			UserNick: user.Nick,
			ParentId: pid,
			Content:  r.FormValue("content"),
		}

		if r.FormValue("resiever") != "" {
			resiever, err := strconv.Atoi(r.FormValue("resiever"))
			if err != nil {
				s.error(w, http.StatusInternalServerError, err)
				return
			}

			comm.Resiever = &resiever
		}

		if err := s.service.Post.AddComm(&comm); err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

	case http.MethodGet:
	default:
		s.error(w, http.StatusMethodNotAllowed, nil)
		return
	}

	post, err := s.service.Post.GetPostById(pid)
	if err != nil {
		s.error(w, http.StatusInternalServerError, err)
		return
	}

	comms, err := s.service.Post.GetCommById(pid)
	if err != nil {
		s.error(w, http.StatusInternalServerError, err)
		return
	}

	data.Post = post
	data.Comms = comms

	tmp, err := template.ParseFiles("./web/static/html/commPost.html")
	if err != nil {
		s.error(w, http.StatusInternalServerError, err)
		return
	}

	if err := tmp.ExecuteTemplate(w, "commPost", data); err != nil {
		s.error(w, http.StatusInternalServerError, err)
		return
	}
}
