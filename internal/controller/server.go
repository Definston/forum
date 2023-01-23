package controller

import (
	"log"
	"net/http"

	"forum/internal/service"
)

type server struct {
	service *service.Service
	handler http.ServeMux
}

func NewServer(serv *service.Service) *server {
	return &server{
		service: serv,
	}
}

func (s *server) Run(addr string) error {
	s.router()
	return http.ListenAndServe(addr, &s.handler)
}

func (s *server) router() {
	s.handler.Handle("/web/", http.StripPrefix("/web", http.FileServer(http.Dir("./web/"))))
	s.handler.HandleFunc("/", s.middleWare(s.homePage))
	s.handler.HandleFunc("/login", s.middleWare(s.login))
	s.handler.HandleFunc("/logout", s.middleWare(s.logout))
	s.handler.HandleFunc("/signup", s.middleWare(s.signup))
	s.handler.HandleFunc("/vote", s.middleWare(s.vote))
	s.handler.HandleFunc("/post", s.middleWare(s.post))
	s.handler.HandleFunc("/post/add", s.middleWare(s.postAdd))
}

func (s *server) middleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := s.service.Sess.GetSession(r)
		if err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		next(w, r.WithContext(ctx))
	}
}

func (s *server) error(w http.ResponseWriter, status int, err error) {
	log.Println(err)
	http.Error(w, http.StatusText(status), status)
	return
}
