package main

import "github.com/gin-gonic/gin"

type server struct {
	router *gin.Engine
	// TODO add logger
	// TODO add store
}

func NewServer() *server {
	s := &server{
		router: gin.Default(),
	}

	s.routes()

	return s
}
func (s *server) Start() {
	if err := s.router.Run(); err != nil {
		panic(err) // TODO use logger's fatal
	}
}
