package middlewares

import (
	"log"
	"net/http"

	"github.com/AppDeveloperMLLB/todo_app/common"
)

type resLoggingWriter struct {
	http.ResponseWriter
	code int
}

func NewResLoggingWriter(w http.ResponseWriter) *resLoggingWriter {
	return &resLoggingWriter{ResponseWriter: w, code: http.StatusOK}
}

func (rsw *resLoggingWriter) WriteHeader(code int) {
	rsw.code = code
	rsw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			traceId := newTraceId()
			log.Printf("[%d]%s %s\n", traceId, req.RequestURI, req.Method)

			ctx := common.SetTraceId(req.Context(), traceId)
			req = req.WithContext(ctx)
			rlw := NewResLoggingWriter(w)

			next.ServeHTTP(rlw, req)

			log.Printf("[%d]res: %d", traceId, rlw.code)
		})
}
