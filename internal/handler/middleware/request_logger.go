package middleware

import (
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func ReqLogger(log *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		log = zap.L().With(
			zap.String("component", "middleware/mw_logger"),
		)

		log.Info("middleware mw_logger enabled")

		fn := func(w http.ResponseWriter, r *http.Request) {
			entry := log.With(
				zap.String("method", r.Method),
				zap.String("path", r.URL.Path),
				zap.String("remote_address", r.RemoteAddr),
				zap.String("user_agent", r.UserAgent()),
				zap.String("request_id", middleware.GetReqID(r.Context())),
				zap.String("ip", r.Header.Get("X-Forwarded-For")),
			)

			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			t := time.Now()

			defer func() {
				entry.Info("request completed",
					zap.Int("status", ww.Status()),
					zap.Int("bytes", ww.BytesWritten()),
					zap.String("duration", time.Since(t).String()),
				)
			}()

			next.ServeHTTP(ww, r)
		}

		return http.HandlerFunc(fn)
	}
}
