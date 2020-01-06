package main

import (
	"database/sql"
	"github.com/elmanelman/sql-judge-api/store"
	"github.com/elmanelman/sql-judge-api/store/postgresstore"
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

	return s
}

func (s *server) Start(config *Config) {
	db, err := sql.Open("postgres", config.PostgresConfig.ConnectionString())
	if err != nil {
		s.logger.Error("unable to connect database:", err)
		return
	}

	s.store = postgresstore.New(db) // TODO make independent of postgres

	s.logger.Level = loggerLevels[config.LoggerConfig.Level]

	if err := s.router.Run(); err != nil {
		s.logger.Fatalf("router error: %s", err)
	}
}
