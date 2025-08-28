package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"atlas/internal/config"
	"atlas/server"
	"atlas/store"
	"atlas/store/db"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	config := &config.Config{
		Addr:        getEnv("ADDR", ""),
		DatabaseURL: getEnv("DATABASE_URL", "atlas.db"),
		Port:        getEnv("PORT", "8080"),
		Environment: getEnv("ENVIRONMENT", "development"),
	}

	driver, err := db.NewDBDriver(config)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	store := store.NewStore(driver, config)

	s, err := server.NewServer(ctx, store, config)
	if err != nil {
		cancel()
		slog.Error("failed to create server", "error", err)
	}

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	if err := s.Start(); err != nil {
		if err != http.ErrServerClosed {
			cancel()
			slog.Error("failed to start server", "error", err)
		}
	}

	address := config.Addr
	if address == "" {
		address = "localhost"
	}
	slog.Info("Server started", "address", address+":"+config.Port)

	go func() {
		<-c
		if err := s.Shutdown(ctx); err != nil {
			slog.Error("failed to shutdown server", "error", err)
		}
		cancel()
	}()

	<-ctx.Done()
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
