package store

import "github.com/elmanelman/sql-judge-api/model"

type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByLogin(string) (*model.User, error)
	//Update(*model.User) error
	//Delete(*model.User) error
}
