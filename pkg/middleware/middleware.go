package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Sulav-Adhikari/gouser/pkg/auth"
)

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		claims, err := auth.ParseToken(tokenString)
		if err != nil {
			http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		// Attach claims to the request context (optional, for further use in route handlers)
		ctx := context.WithValue(r.Context(), "claims", claims)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
