package model

type User struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
	About    string `json:"about"`
	RoleID   int    `json:"role_id"`
}
