package httpserver

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewHTTPServer(
	httpAddr string,
) *http.Server {
	router := chi.NewRouter()

	router.Use(RequestLogger(slog.Default()))

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		if err := WriteJSON(w, http.StatusOK, map[string]string{"status": "ok"}); err != nil {
			slog.Error("write health response failed", "error", err)
		}
	})

	router.Get("/error", func(w http.ResponseWriter, r *http.Request) {
		WriteMappedError(w, ErrForbidden)
	})

	return &http.Server{
		Addr:    httpAddr,
		Handler: router,
	}
}
