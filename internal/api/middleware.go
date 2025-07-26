package api

import (
	"log"
	"net/http"
	"time"
	"strings"
	"context"
	
	"github.com/thesaltysleuth/notes-service/internal/auth"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w,r)

		log.Printf("%s %s took %v", r.Method, r.URL.Path, time.Since(start))
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer") {
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		claims,err := auth.ValidateToken(tokenStr)
		if err!= nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
    // no lint
		type username struct{}
		ctx := context.WithValue(r.Context(), username{}, claims.Username)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
