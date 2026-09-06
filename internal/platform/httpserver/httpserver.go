package httpserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewHTTPServer(
	httpAddr string,
) *http.Server {
	router := chi.NewRouter()

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {

		if err := WriteJSON(w, http.StatusOK, map[string]string{"status": "ok"}); err != nil {
			// serverseitig loggen; keine zweite Response schreiben
		}
	})

	router.Get("/error", func(w http.ResponseWriter, r *http.Request) {
		WriteError(w, http.StatusBadRequest, "ABC111", "TEST ERROR")
	})

	return &http.Server{
		Addr:    httpAddr,
		Handler: router,
	}
}
