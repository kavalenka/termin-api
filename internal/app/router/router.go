package router

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"termin-api/internal/app/resource/health"
	"termin-api/internal/app/resource/user"
)

func New(db *sql.DB, validator *validator.Validate) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", health.Get)

	userApi := user.New(db, validator)
	r.Get("/users", userApi.List)
	r.Post("/users", userApi.Create)

	return r
}
