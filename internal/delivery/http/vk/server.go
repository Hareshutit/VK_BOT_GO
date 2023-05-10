package vk

import (
	"net/http"

	"VK_BOT_GO/internal/usecase"
)

type Server struct {
	Сommand usecase.Command
	Server  http.Server
}

func New() *Server {
	return &Server{
		Сommand: usecase.Command{},
		Server:  http.Server{},
	}
}

func (s *Server) Check(str string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if http.MethodPost != r.Method {
			http.NotFound(w, r)
			return
		}
		w.Write([]byte(str))
	}
}
