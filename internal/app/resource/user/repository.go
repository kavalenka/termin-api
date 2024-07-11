package user

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) List() (Users, error) {
	query := "select * from users;"
	results, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	return listUsers(results)
}

func (r *Repository) FindByEmailOrPhone(user *User) (Users, error) {
	query := "select * from users where email=$1 or phone=$2"
	results, err := r.db.Query(query, user.Email, user.Phone)

	if err != nil {
		return nil, err
	}

	return listUsers(results)
}

func (r *Repository) Create(user *User) error {
	query := "insert into users (name, email, phone, password, dob, role, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7, $8);"
	_, err := r.db.Exec(query, user.Name, user.Email, user.Phone, user.Password, user.DOB, user.Role, user.CreatedAt, user.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

func listUsers(results *sql.Rows) (Users, error) {
	users := make([]*User, 0)

	for results.Next() {
		var u User
		results.Scan(&u.ID, &u.Name, &u.Email, &u.Phone, &u.Password, &u.DOB, &u.Role, &u.CreatedAt, &u.UpdatedAt)
		users = append(users, &u)
	}

	return users, nil
}
