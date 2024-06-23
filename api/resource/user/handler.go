package user

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type API struct {
	repository *Repository
}

func New(db *sql.DB) *API {
	return &API{
		repository: NewRepository(db),
	}
}

func (a *API) List(w http.ResponseWriter, _ *http.Request) {
	users, _ := a.repository.List()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users.Serialized())
}
