package main

import (
	"github.com/elmanelman/sql-judge-api/store"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type server struct {
	router *gin.Engine
	logger *logrus.Logger
	store  store.Store
}

func NewServer(store store.Store) *server {
	s := &server{
		router: gin.Default(),
		logger: logrus.New(),
		store:  store,
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
