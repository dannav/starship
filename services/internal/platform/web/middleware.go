package web

import (
	"bufio"
	"errors"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/pborman/uuid"
)

const requestIDHeader = "X-Request-Id"

// responseWriter wraps an http.ResponseWriter so we can
// capture the status code.
type responseWriter struct {
	status int
	http.ResponseWriter
}

// WriteHeader captures the statusCode and then writes it the
// wrapped ResponseWriter.
func (w *responseWriter) WriteHeader(statusCode int) {
	w.status = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

// Hijack implements the http.Hijacker interface.
func (w *responseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	h, ok := w.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, errors.New("ResponseWriter does not implement http.Hijacker")
	}
	return h.Hijack()
}

// RequestMW is a middleware that creates a request id for each request
// and sets it on the header field X-Request-Id. Also logs the start and
// end of each request.
func RequestMW(next http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {

		st := time.Now()

		ww := &responseWriter{
			status:         http.StatusOK,
			ResponseWriter: w,
		}

		// Check if request ID was passed in header. Otherwise, generate one.
		id := r.Header.Get("X-Request-Id")
		if id == "" {
			id = uuid.New()
		}

		defer func() {
			log.Printf("%s %s - complete [%s] %d - %s", r.Method, r.RequestURI, time.Since(st), ww.status, id)
		}()

		ww.Header().Set(requestIDHeader, id)

		next.ServeHTTP(ww, r)
	}
	return http.HandlerFunc(f)
}
