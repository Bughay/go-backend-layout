// internal/auth/middleware.go
package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Bughay/go-backend-layout/internal/model"
)

// contextKey is an unexported type to prevent context key collisions.
type contextKey string

const claimsContextKey contextKey = "jwt_claims"

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, message string) {
	WriteJSON(w, status, model.APIError{Code: status, Message: message})
}

// Middleware returns an http.Handler middleware that validates the JWT on every request.
func (m *Manager) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			WriteError(w, http.StatusUnauthorized, "authorization header is required")
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "bearer") {
			WriteError(w, http.StatusUnauthorized, "authorization header format must be: Bearer {token}")
			return
		}

		claims, err := m.Validate(parts[1])
		if err != nil {
			WriteError(w, http.StatusUnauthorized, err.Error())
			return
		}

		// Inject the validated claims into the request context
		ctx := context.WithValue(r.Context(), claimsContextKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// ClaimsFromContext retrieves the validated JWT claims from the request context.
// Returns nil if not present (i.e., called on an unprotected route).
func ClaimsFromContext(ctx context.Context) *model.Claims {
	claims, _ := ctx.Value(claimsContextKey).(*model.Claims)
	return claims
}

// RequireRole returns a middleware that enforces a specific user role.
func (m *Manager) RequireRole(role string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims := ClaimsFromContext(r.Context())
		if claims == nil || claims.Role != role {
			WriteError(w, http.StatusForbidden, "you do not have permission to perform this action")
			return
		}
		next.ServeHTTP(w, r)
	})
}
