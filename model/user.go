package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
	About    string `json:"about,omitempty"`
	RoleID   int    `json:"role_id,omitempty"`
}

func (u *User) HashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	u.Password = string(hash)
	return nil
}

func (u *User) Sanitize() {
	u.Password = ""
}
