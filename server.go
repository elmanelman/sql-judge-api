package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type server struct {
	router *gin.Engine
	logger *logrus.Logger
	// TODO add store
}

func NewServer() *server {
	s := &server{
		router: gin.Default(),
		logger: logrus.New(),
	}

	s.routes()

	s.configureLogger()

	return s
}

func (s *server) configureLogger() {
	s.logger.SetLevel(logrus.DebugLevel)
}

func (s *server) Start() {
	if err := s.router.Run(); err != nil {
		s.logger.Fatalf("router error: %s", err)
	}
}
