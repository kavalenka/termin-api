package router

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"termin-api/internal/app/resource/health"
	"termin-api/internal/app/resource/user"
)

func New(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", health.Get)

	userApi := user.New(db)
	r.Get("/users", userApi.List)
	r.Post("/users", userApi.Create)

	return r
}
