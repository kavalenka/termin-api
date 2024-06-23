package user

import (
	"database/sql"
	"log"

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
	users := make([]*User, 0)
	query := "select * from users;"
	results, err := r.db.Query(query)

	if err != nil {
		log.Fatalln("Error while fetching users:", err)
	}

	for results.Next() {
		var u User
		results.Scan(&u.ID, &u.Name, &u.Email, &u.Phone, &u.Password, &u.DOB, &u.Role, &u.CreatedAt, &u.UpdatedAt)
		users = append(users, &u)
	}

	return users, nil
}
