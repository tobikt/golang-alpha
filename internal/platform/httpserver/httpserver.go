package httpserver

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewHTTPServer(
	httpAddr string,
) *http.Server {
	router := chi.NewRouter()

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":"ok"}`)
	})

	return &http.Server{
		Addr:    httpAddr,
		Handler: router,
	}
}
