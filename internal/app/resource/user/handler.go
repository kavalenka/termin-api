package user

import (
	"database/sql"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"

	"termin-api/internal/app/resource/common/error"
)

type API struct {
	repository *Repository
	validator  *validator.Validate
}

func New(db *sql.DB, validator *validator.Validate) *API {
	return &API{
		repository: NewRepository(db),
		validator:  validator,
	}
}

func (a *API) List(w http.ResponseWriter, _ *http.Request) {
	users, err := a.repository.List()

	if err != nil {
		error.InternalServerError(w, error.DBDataReadFailure)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users.Serialized())
}

func (a *API) Create(w http.ResponseWriter, r *http.Request) {
	createParams := &CreateParams{}

	json.NewDecoder(r.Body).Decode(createParams)

	err := a.validator.Struct(createParams)
	if err != nil {
		error.ValidationError(w, error.ValidationFailure)
		return
	}

	newUser := createParams.ToUser()

	existingUsers, err := a.repository.FindByEmailOrPhone(newUser)

	if err != nil {
		error.InternalServerError(w, error.DBDataReadFailure)
		return
	}

	if len(existingUsers) > 0 {
		error.UnprocessableEntity(w, error.EmailOrPhoneExists)
		return
	}

	err = a.repository.Create(newUser)

	if err != nil {
		error.InternalServerError(w, error.DBDataInsertFailure)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
