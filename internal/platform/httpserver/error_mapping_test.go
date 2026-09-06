package httpserver

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWriteMappedErrorRecognizesWrappedNotFound(t *testing.T) {
	recorder := httptest.NewRecorder()

	err := fmt.Errorf("lookup failed: %w", ErrNotFound)

	WriteMappedError(recorder, err)

	if recorder.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", recorder.Code)
	}

	if strings.Contains(recorder.Body.String(), "lookup failed") {
		t.Fatal("internal error details must not be exposed")
	}
}

func TestWriteMappedErrorRecognizesWrappedInvalidRequest(t *testing.T) {
	recorder := httptest.NewRecorder()

	err := fmt.Errorf("lookup failed: %w", ErrInvalidRequest)

	WriteMappedError(recorder, err)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", recorder.Code)
	}

	if strings.Contains(recorder.Body.String(), "lookup failed") {
		t.Fatal("internal error details must not be exposed")
	}
}

func TestWriteMappedErrorRecognizesWrappedForbidden(t *testing.T) {
	recorder := httptest.NewRecorder()

	err := fmt.Errorf("lookup failed: %w", ErrForbidden)

	WriteMappedError(recorder, err)

	if recorder.Code != http.StatusForbidden {
		t.Fatalf("expected status 403, got %d", recorder.Code)
	}

	if strings.Contains(recorder.Body.String(), "lookup failed") {
		t.Fatal("internal error details must not be exposed")
	}
}

func TestWriteMappedErrorRecognizesWrappedInternalServerError(t *testing.T) {
	recorder := httptest.NewRecorder()

	err := fmt.Errorf("lookup failed: %w", ErrInternalError)

	WriteMappedError(recorder, err)

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 500, got %d", recorder.Code)
	}

	if strings.Contains(recorder.Body.String(), "lookup failed") {
		t.Fatal("internal error details must not be exposed")
	}
}

func TestWriteMappedErrorHidesUnknownError(t *testing.T) {
	recorder := httptest.NewRecorder()

	WriteMappedError(recorder, errors.New("database password leaked"))

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 500, got %d", recorder.Code)
	}

	if strings.Contains(recorder.Body.String(), "database password leaked") {
		t.Fatal("internal error details must not be exposed")
	}
}
