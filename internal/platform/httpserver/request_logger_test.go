package httpserver

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestLogger(t *testing.T) {
	var logBuffer bytes.Buffer

	logger := slog.New(
		slog.NewJSONHandler(&logBuffer, nil),
	)

	var handlerRequestID string
	handlerCalls := 0

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlerCalls++

		handlerRequestID = RequestID(r.Context())

		w.WriteHeader(http.StatusTeapot)
	})

	handler := RequestLogger(logger)(next)

	request := httptest.NewRequest(
		http.MethodGet,
		"/health",
		nil,
	)
	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, request)

	if handlerCalls != 1 {
		t.Fatalf("expected handler to be called once, got %d", handlerCalls)
	}

	responseRequestID := recorder.Header().Get("X-Request-ID")

	if responseRequestID == "" {
		t.Fatal("expected X-Request-ID response header")
	}

	if handlerRequestID == "" {
		t.Fatal("expected request ID in handler context")
	}

	if handlerRequestID != responseRequestID {
		t.Fatalf(
			"expected context request ID %q, got response header %q",
			handlerRequestID,
			responseRequestID,
		)
	}

	if recorder.Code != http.StatusTeapot {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusTeapot,
			recorder.Code,
		)
	}

	var logEntry map[string]any

	if err := json.Unmarshal(logBuffer.Bytes(), &logEntry); err != nil {
		t.Fatalf("log is not valid JSON: %v", err)
	}

	assertLogValue(t, logEntry, "method", http.MethodGet)
	assertLogValue(t, logEntry, "path", "/health")
	assertLogValue(t, logEntry, "status", float64(http.StatusTeapot))
	assertLogValue(t, logEntry, "request_id", responseRequestID)

	if _, exists := logEntry["duration"]; !exists {
		t.Fatal("expected duration in log entry")
	}
}

func TestRequestLoggerDefaultsToOK(t *testing.T) {
	var logBuffer bytes.Buffer
	logger := slog.New(slog.NewJSONHandler(&logBuffer, nil))

	next := http.HandlerFunc(func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		_, _ = w.Write([]byte("ok"))
	})

	handler := RequestLogger(logger)(next)

	request := httptest.NewRequest(http.MethodGet, "/health", nil)
	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", recorder.Code)
	}

	var logEntry map[string]any
	if err := json.Unmarshal(logBuffer.Bytes(), &logEntry); err != nil {
		t.Fatalf("log is not valid JSON: %v", err)
	}

	assertLogValue(t, logEntry, "status", float64(http.StatusOK))
}

func assertLogValue(t *testing.T, entry map[string]any, key string, want any) {
	t.Helper()

	got, exists := entry[key]
	if !exists {
		t.Fatalf("expected log field %q", key)
	}

	if got != want {
		t.Fatalf("expected log field %q to be %v, got %v", key, want, got)
	}
}
