package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func ServeRoutes(r *chi.Mux) {
	Routes(r)
}

func getField(r *http.Request, name string) string {
	field := chi.URLParam(r, name)
	return field
}