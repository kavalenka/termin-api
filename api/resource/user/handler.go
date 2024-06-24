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

func (a *API) Create(w http.ResponseWriter, r *http.Request) {
	createParams := &CreateParams{}

	json.NewDecoder(r.Body).Decode(createParams)
	newUser := createParams.ToUser()
	err := a.repository.Create(newUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
