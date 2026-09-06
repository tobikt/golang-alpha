package httpserver

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"log/slog"
	"net/http"
	"time"
)

type requestIDKey struct{}

func RequestID(ctx context.Context) string {
	requestID, _ := ctx.Value(requestIDKey{}).(string)
	return requestID
}

func setRequestId(r *http.Request, requestID string) *http.Request {
	ctx := context.WithValue(
		r.Context(),
		requestIDKey{},
		requestID,
	)
	r = r.WithContext(ctx)
	return r
}

func RequestLogger(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			requestID := newRequestID()

			// Writes the requestID into the context with requestIDKey
			r = setRequestId(r, requestID)

			w.Header().Set("X-Request-ID", requestID)

			writer := &statusWriter{
				ResponseWriter: w,
			}

			next.ServeHTTP(writer, r)

			logger.Info(
				"request",
				"method", r.Method,
				"path", r.URL.Path,
				"status", writer.status,
				"duration", time.Since(start),
				"request_id", requestID,
			)
		})
	}
}

func newRequestID() string {
	randomBytes := make([]byte, 16)

	if _, err := rand.Read(randomBytes); err != nil {
		return "unknown"
	}

	return hex.EncodeToString(randomBytes)
}

type statusWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *statusWriter) Write(body []byte) (int, error) {
	if w.status == 0 {
		w.WriteHeader(http.StatusOK)
	}

	return w.ResponseWriter.Write(body)
}
