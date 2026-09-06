package httpserver

import (
	"errors"
	"net/http"
)

var ErrNotFound = errors.New("resource not found")
var ErrInvalidRequest = errors.New("invalid request")
var ErrForbidden = errors.New("forbidden")
var ErrInternalError = errors.New("internal server error")

func WriteMappedError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, ErrNotFound):
		WriteError(w, http.StatusNotFound, "resource_not_found", "resource not found")
	case errors.Is(err, ErrInvalidRequest):
		WriteError(w, http.StatusBadRequest, "invalid_request", "invalid request")
	case errors.Is(err, ErrForbidden):
		WriteError(w, http.StatusForbidden, "forbidden", "forbidden")
	case errors.Is(err, ErrInternalError):
		WriteError(w, http.StatusInternalServerError, "internal_error", "an internal error occurred")
	default:
		WriteError(w, http.StatusInternalServerError, "internal_error", "an internal error occurred")
	}
}
