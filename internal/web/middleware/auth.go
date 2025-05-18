package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/sahay-shashank/mongodb-server/internal/core/auth"
	context_keys "github.com/sahay-shashank/mongodb-server/internal/core/context"
)

func jwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Missing or invalid Authorization header", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := auth.ValidateJWT(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Store tenant info in context
		ctx := context.WithValue(r.Context(), context_keys.TenantIDKey, claims.TenantID)
		ctx = context.WithValue(ctx, context_keys.ServiceKey, claims.Service)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
