package main

import (
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"suggestions/internal/app/router"
	"suggestions/internal/config"
	"suggestions/internal/logger"
	"suggestions/internal/usecases"
	"suggestions/pkg/suggestions"
	"syscall"
)

func main() {
	log := logger.New(slog.LevelInfo)

	cfg, err := config.New()
	if err != nil {
		log.Error("cannot load config", err)
		os.Exit(1)
	}

	sugg := suggestions.NewSuggestions(cfg.YandexConfig.ApiKey, cfg.YandexConfig.URL)
	uc := usecases.NewUsecases(sugg)

	srv := grpc.NewServer()
	_ = router.NewGRPCServer(srv, uc, log)

	listener, err := net.Listen("tcp", ":"+cfg.GRPCConfig.Port)
	if err != nil {
		log.Error("failed to listen address", err)
		os.Exit(1)
	}
	go func() {
		if err = srv.Serve(listener); err != nil {
			log.Error("failed to start gRPC server", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	log.Info("application started")
	select {
	case q := <-quit:
		log.Info("exit from app ", q)
	}
}
