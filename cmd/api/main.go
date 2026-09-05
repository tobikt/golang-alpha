package main

import (
	"context"
	"errors"
	"log"
	"net/http"

	"kuekelheim.de/golang-alpha/internal/platform/config"
	"kuekelheim.de/golang-alpha/internal/platform/database"
	"kuekelheim.de/golang-alpha/internal/platform/httpserver"
)

func main() {
	ctx := context.Background()

	cfg := config.Load()

	db, err := database.NewPostgres(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	server := httpserver.NewHTTPServer(cfg.HTTPAddr)
	log.Printf("API listening on %s", cfg.HTTPAddr)
	if err := server.ListenAndServe(); err != nil &&
		!errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}

}
