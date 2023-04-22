package server

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ettore83/projeto-anime/anime"
	"github.com/ettore83/projeto-anime/internal/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Router struct {
	Service *service.Service
}

// NewRouter Return a new basic router with some handy middleware
func NewRouter(ctx context.Context, svc *service.Service) chi.Router {
	r := chi.NewRouter()
	router := Router{Service: svc}

	r.Use(
		middleware.SetHeader("Content-Type", "application/json"),
	)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusMethodNotAllowed)
	})

	r.Route("/", router.routes(ctx))
	return r
}

func (ro Router) routes(ctx context.Context) func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/", ro.liveness())
		r.Post("/anime", ro.insert(ctx))
		r.Get("/anime", ro.list(ctx))
	}
}

func (r Router) liveness() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		_, _ = res.Write([]byte(`["Server is live"]`))
	}
}

func (r Router) insert(ctx context.Context) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		anime := &anime.Anime{}
		json.Unmarshal(b, anime)

		err = r.Service.Insert(ctx, anime)
		if err != nil {
			log.Printf("Error %v", err)
			res.WriteHeader(http.StatusInternalServerError)
			return
		}

		res.WriteHeader(http.StatusOK)
		r, _ := json.Marshal(anime)
		_, _ = res.Write(r)
	}
}

func (r Router) list(ctx context.Context) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		animes, err := r.Service.List(ctx)
		if err != nil {
			log.Printf("Error %v", err)
			res.WriteHeader(http.StatusInternalServerError)
			return
		}

		res.WriteHeader(http.StatusOK)
		r, _ := json.Marshal(animes)
		_, _ = res.Write(r)
	}
}
