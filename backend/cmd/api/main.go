package main

import (
	"context"
	"log"

	"github.com/ettore83/projeto-anime/cmd/api/internal/config"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("config", zap.Error(err))
	}

	if err = run(ctx, cfg); err != nil {
		log.Fatal("command failed", zap.Error(err))
	}
	log.Print("command completed successfully")
}
