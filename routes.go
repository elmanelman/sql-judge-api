package main

func (s *server) routes() {
	// create user
	s.router.POST("/user", s.handleUserCreate())
}
