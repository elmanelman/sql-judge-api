package postgresstore

import (
	"github.com/elmanelman/sql-judge-api/model"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.HashPassword(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO user_account(login, password) VALUES ($1, $2) RETURNING ID",
		u.Login,
		u.Password,
	).Scan(&u.ID)
}

func (r *UserRepository) Find(id int) (*model.User, error) {
	u := &model.User{}

	if err := r.store.db.QueryRow(
		"SELECT id, login, password, COALESCE(email, ''), COALESCE(about, ''), COALESCE(role_id, 0) FROM user_account WHERE id = $1",
		id,
	).Scan(
		&u.ID,
		&u.Login,
		&u.Password,
		&u.Email,
		&u.About,
		&u.RoleID,
	); err != nil {
		return nil, err
	}

	u.Sanitize()

	return u, nil
}

func (r *UserRepository) FindByLogin(login string) (*model.User, error) {
	u := &model.User{}

	if err := r.store.db.QueryRow(
		"SELECT id, login, password, COALESCE(email, ''), COALESCE(about, ''), COALESCE(role_id, 0) FROM user_account WHERE login = $1",
		login,
	).Scan(
		&u.ID,
		&u.Login,
		&u.Password,
		&u.Email,
		&u.About,
		&u.RoleID,
	); err != nil {
		return nil, err
	}

	u.Sanitize()

	return u, nil
}
