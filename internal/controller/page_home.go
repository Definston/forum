package controller

import (
	"html/template"
	"net/http"

	"forum/internal/model"
)

func (s *server) homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		s.error(w, http.StatusNotFound, nil)
		return
	}

	uid := r.Context().Value(model.ContextKeySession).(int)

	data := struct {
		User  *model.User
		Posts *[]model.Post
	}{}

	user, err := s.service.Auth.GetUserById(uid)
	if err != nil {
		s.error(w, http.StatusInternalServerError, err)
		return
	}

	data.User = user

	switch r.Method {
	case http.MethodGet:
		if r.URL.Query().Has("filter") {

			filter := r.URL.Query().Get("filter")

			switch filter {
			case "liked":
				if uid <= 0 {
					s.error(w, http.StatusUnauthorized, nil)
					return
				}

				posts, err := s.service.Post.GetPostsLikedByUser(uid)
				if err != nil {
					s.error(w, http.StatusInternalServerError, err)
					return
				}

				data.Posts = posts

			case "users":
				if uid <= 0 {
					s.error(w, http.StatusUnauthorized, nil)
					return
				}

				posts, err := s.service.Post.GetPostsAddedByUser(uid)
				if err != nil {
					s.error(w, http.StatusInternalServerError, err)
					return
				}

				data.Posts = posts

			default:
				if len(filter) == 0 {
					s.error(w, http.StatusBadRequest, nil)
					return
				}

				posts, err := s.service.Post.GetPostsByTag(filter)
				if err != nil {
					s.error(w, http.StatusInternalServerError, err)
					return
				}

				data.Posts = posts

			}

		} else if len(r.URL.Query()) != 0 {

			s.error(w, http.StatusBadRequest, nil)
			return

		} else {

			posts, err := s.service.Post.GetAllPosts()
			if err != nil {
				s.error(w, http.StatusInternalServerError, err)
				return

			}

			data.Posts = posts
		}

	default:
		s.error(w, http.StatusMethodNotAllowed, nil)
		return
	}

	tmp, err := template.ParseFiles("./web/static/html/homePage.html")
	if err != nil {
		s.error(w, http.StatusInternalServerError, err)
		return
	}

	if err := tmp.ExecuteTemplate(w, "home", data); err != nil {
		s.error(w, http.StatusInternalServerError, err)
		return
	}
}
