package router

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"termin-api/api/resource/health"
	"termin-api/api/resource/user"
)

func New(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", health.Get)

	userApi := user.New(db)
	r.Get("/users", userApi.List)

	return r
}
