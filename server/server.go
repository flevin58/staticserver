package server

import (
	"errors"
	"log"
	"net/http"
)

type Server struct {
	address string
	router  *Router
}

func New() *Server {
	server := &Server{
		router:  NewRouter(),
		address: ":8080",
	}
	return server
}

func (s *Server) WithAddress(address string) *Server {
	s.address = address
	return s
}

func (s *Server) WithStaticDir(folder string) *Server {
	s.router.defaultHandler = http.FileServer(http.Dir(folder))
	return s
}

func (s *Server) WithRoute(path, method string, handler http.HandlerFunc) *Server {
	s.AddRoute(path, method, handler)
	return s
}

func (s *Server) Run() {
	log.Println("server listening for connections on address ", s.address)
	if err := http.ListenAndServe(s.address, s.router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			log.Println("server closed")
		}
		log.Fatalf("server failure: %v", err)
	}
}

func (s *Server) AddRoute(path, method string, handler http.HandlerFunc) {
	s.router.AddRoute(path, method, handler)
}
