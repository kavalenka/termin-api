package router

import (
	"github.com/go-chi/chi/v5"
	"termin-api/api/resource/health"
)

func New() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", health.Get)

	return r
}
