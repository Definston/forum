package controller

import (
	"net/http"
	"text/template"

	"forum/internal/model"
)

func (s *server) postAdd(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/post/add" {
		s.error(w, http.StatusMethodNotAllowed, nil)
		return
	}

	uid := r.Context().Value(model.ContextKeySession).(int)
	if uid <= 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	user, err := s.service.Auth.GetUserById(uid)
	if err != nil {
		s.error(w, http.StatusInternalServerError, err)
		return
	}

	if user == nil {
		s.error(w, http.StatusUnauthorized, err)
		return
	}

	switch r.Method {
	case http.MethodPost:

		if !s.service.Post.IsValidContent(r.FormValue("content")) {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		tag, err := s.service.Tag.FindTags(r.FormValue("content"))
		if err != nil {
			s.error(w, http.StatusMethodNotAllowed, err)
			return
		}

		post := &model.Post{
			UserId:   user.Id,
			UserNick: user.Nick,
			Tag:      *tag,
			Content:  r.FormValue("content"),
		}

		if err := s.service.Post.AddPost(post); err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		http.Redirect(w, r, "/", http.StatusFound)
		return

	case http.MethodGet:
	default:
		s.error(w, http.StatusMethodNotAllowed, nil)
		return
	}

	tmp, err := template.ParseFiles("./web/static/html/addPost.html")
	if err != nil {
		s.error(w, http.StatusInternalServerError, err)
		return
	}

	if err := tmp.ExecuteTemplate(w, "addPost", nil); err != nil {
		s.error(w, http.StatusInternalServerError, err)
		return
	}
}
