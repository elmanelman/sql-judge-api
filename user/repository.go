package user

import "github.com/elmanelman/sql-judge-api/model"

type Repository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByLogin(string) (*model.User, error)
	Update(*model.User) error
	Delete(*model.User) error
}
