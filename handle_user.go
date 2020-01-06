package main

import (
	"github.com/elmanelman/sql-judge-api/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *server) handleUserCreate() gin.HandlerFunc {
	type request struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	type response struct {
		ID    int    `json:"id"`
		Login string `json:"login"`
	}

	return func(c *gin.Context) {
		req := &request{}
		if err := c.BindJSON(req); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		u := &model.User{
			Login:    req.Login,
			Password: req.Password,
		}
		if err := s.store.User().Create(u); err != nil {
			c.AbortWithStatus(http.StatusUnprocessableEntity)
			return
		}

		rsp := &response{
			ID:    u.ID,
			Login: u.Login,
		}
		c.JSON(http.StatusOK, rsp)
	}
}
