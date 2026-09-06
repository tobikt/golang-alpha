package httpserver

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJsonWriter(t *testing.T) {
	recorder := httptest.NewRecorder()

	err := WriteJSON(
		recorder,
		http.StatusCreated,
		map[string]string{"name": "Test"},
	)
	if err != nil {
		t.Fatalf("WriteJson returned an error: %v", err)
	}

	if recorder.Code != http.StatusCreated {
		t.Fatalf("excepted status %d, got %d",
			http.StatusCreated,
			recorder.Code,
		)
	}

	if got := recorder.Header().Get("Content-Type"); got != "application/json" {
		t.Fatalf("expected Content-Type application/json, got %q", got)
	}

	var body map[string]string
	if err := json.Unmarshal(recorder.Body.Bytes(), &body); err != nil {
		t.Fatalf("response body is not valid JSON: %v", err)
	}

	if body["name"] != "Test" {
		t.Fatalf("expected name Test, got %q", body["name"])
	}
}

func TestWriteError(t *testing.T) {
	recorder := httptest.NewRecorder()

	WriteError(
		recorder,
		http.StatusBadRequest,
		"XXXABC",
		"TEST ERROR",
	)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("excepted status %d, got %d",
			http.StatusBadRequest,
			recorder.Code,
		)
	}

	var body errorResponse
	if err := json.Unmarshal(recorder.Body.Bytes(), &body); err != nil {
		t.Fatalf("response body is not valid JSON: %v", err)
	}

	if body.Error.Code != "XXXABC" {
		t.Fatalf("expected error code invalid_request, got %q",
			body.Error.Code)
	}

	if body.Error.Message != "TEST ERROR" {
		t.Fatalf("unexpected error message: %q",
			body.Error.Message)
	}
}

func TestWriteJSONReturnsErrorBeforeWritingResponse(t *testing.T) {
	recorder := httptest.NewRecorder()

	err := WriteJSON(
		recorder,
		http.StatusOK,
		make(chan int),
	)
	if err == nil {
		t.Fatal("expected JSON encoding error, got nil")
	}

	if recorder.Body.Len() != 0 {
		t.Fatalf("expected empty response body, got %q",
			recorder.Body.String())
	}
}
