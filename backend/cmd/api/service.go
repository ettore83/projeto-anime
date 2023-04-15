package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/ettore83/projeto-anime/cmd/api/internal/config"
	"github.com/ettore83/projeto-anime/internal/server"
)

func run(ctx context.Context, cfg *config.Config) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	r := server.NewRouter()
	srv := server.New(r, cfg.Http.Address, cfg.Http.WriteTimeout, cfg.Http.ReadTimeout, cfg.Http.IdleTimeout)

	done := make(chan struct{}, 1)
	go func(done chan<- struct{}) {
		<-ctx.Done()

		log.Print("Stopping the server...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		srv.SetKeepAlivesEnabled(false)
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal("Could not gracefully shutdown the server.")
		}

		log.Print("Server stopped.")
		close(done)
	}(done)

	log.Print("Server running...")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	<-done

	return nil
}
