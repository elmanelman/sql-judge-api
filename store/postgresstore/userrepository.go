package postgresstore

import "github.com/elmanelman/sql-judge-api/model"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) error {
	return r.store.db.QueryRow(
		"INSERT INTO user_account(login, password) VALUES ($1, $2) RETURNING ID",
		u.Login,
		u.Password,
	).Scan(&u.ID)
}
