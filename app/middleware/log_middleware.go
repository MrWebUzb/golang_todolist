package middleware

import (
	"net/http"
	"time"

	"github.com/MrWebUzb/apiserver/app/constants"

	"github.com/sirupsen/logrus"
)

// WrappedResponseWriter response writer for logging requests
// Customize default response writer
type WrappedResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

// NewWrappedResponseWriter constructor function
func NewWrappedResponseWriter(rw http.ResponseWriter) *WrappedResponseWriter {
	return &WrappedResponseWriter{rw, http.StatusOK}
}

// Header this function overrides default ResponseWriter.Header function
func (rw *WrappedResponseWriter) Header() http.Header {
	return rw.ResponseWriter.Header()
}

// Write this function overrides default ResponseWriter.Write function
func (rw *WrappedResponseWriter) Write(data []byte) (int, error) {
	return rw.ResponseWriter.Write(data)
}

// WriteHeader this function overrides default ResponseWriter.WriteHeader function
func (rw *WrappedResponseWriter) WriteHeader(statusCode int) {
	rw.StatusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

// RequestLog middleware for catching requests
func RequestLog(log *logrus.Logger, router http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		// Set content type to application/json
		rw.Header().Add("Content-Type", "application/json; charset=utf-8")
		wrw := NewWrappedResponseWriter(rw)
		// Start time of serving request
		start := time.Now()
		router.ServeHTTP(wrw, r)

		logByRequest(
			log,
			wrw.StatusCode,
			r.URL.Path,
			time.Since(start),
		)
	})
}

func logByRequest(log *logrus.Logger, sc int, path string, exec time.Duration) {
	switch sc {
	case http.StatusOK:
		log.Infof("%s[%d-%s] %s%s - %.4f ms",
			constants.ConsoleColorGreen,
			http.StatusOK,
			http.StatusText(sc),
			path,
			constants.ConsoleColorReset,
			nanoToMilli(exec),
		)
	case http.StatusNotFound:
		log.Warnf("%s[%d-%s] %s%s - %.4f ms",
			constants.ConsoleColorYellow,
			sc,
			http.StatusText(sc),
			path,
			constants.ConsoleColorReset,
			nanoToMilli(exec),
		)
	case http.StatusBadRequest:
		log.Warnf("%s[%d-%s] %s%s - %.4f ms",
			constants.ConsoleColorYellow,
			sc,
			http.StatusText(sc),
			path,
			constants.ConsoleColorReset,
			nanoToMilli(exec),
		)
	case http.StatusMethodNotAllowed:
		log.Warnf("%s[%d-%s] %s%s - %.4f ms",
			constants.ConsoleColorYellow,
			sc,
			http.StatusText(sc),
			path,
			constants.ConsoleColorReset,
			nanoToMilli(exec),
		)
	default:
		log.Errorf("%s[%d-%s] %s%s - %.4f ms",
			constants.ConsoleColorRed,
			sc,
			http.StatusText(sc),
			path,
			constants.ConsoleColorReset,
			nanoToMilli(exec),
		)
	}
}

func nanoToMilli(t time.Duration) float64 {
	return float64(t) / float64(time.Millisecond)
}
