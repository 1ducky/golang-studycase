package middleware

import (
	"log"
	"net/http"
	"restApi/internal/auth"
	httpHelper "restApi/internal/http"
	"restApi/internal/logging"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		Role := "anymous"
		user, ok := auth.GetAuht(r.Context())
		if ok {
			Role = string(user.Role)
		}
		rw := &httpHelper.ResponseWriter{ResponseWriter: w, StatusCode: http.StatusOK}

		requestId := logging.GenerateRandomID()

		ctx := logging.WithRequestID(r.Context(), requestId)
		next.ServeHTTP(rw, r.WithContext(ctx))
		endTime := time.Since(startTime)
		log.Print("Request: ", requestId, " [", rw.StatusCode, "]  ", r.URL.Path, " ", r.Method, " ", endTime, ", Client: ", Role)
	})
}
