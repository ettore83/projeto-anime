package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Router struct {
}

// NewRouter Return a new basic router with some handy middleware
func NewRouter() chi.Router {
	r := chi.NewRouter()

	r.Use(
		middleware.SetHeader("Content-Type", "application/json"),
	)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusMethodNotAllowed)
	})

	r.Route("/", routes())
	return r
}

func routes() func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/", liveness())
	}
}

func liveness() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		_, _ = res.Write([]byte(`["Server is live"]`))
	}
}
