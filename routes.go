package main

func (s *server) routes() {
	// create user
	s.router.POST("/user", s.handleUserCreate())

	// find user by id
	s.router.GET("/user/id/:id", s.handleUserFind())

	// find user by login
	s.router.GET("/user/login/:login", s.handleUserFindByLogin())
}
