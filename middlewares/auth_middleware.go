package middlewares

import (
	"context"
	"net/http"

	"backend/models"
	"backend/utils"
)

type contextKey string

const UserClaimsKey contextKey = "userClaims"

func NewAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("access_token")
		if err != nil {
			utils.WriteJsonErrors(w, models.ErrorJson{Status: 401, Error: "Missing Authorization header"})
			return
		}

		claims, err := utils.VerifyJWT(cookie.Value)
		if err != nil {
			utils.WriteJsonErrors(w, models.ErrorJson{Status: 401, Error: "Invalid or expired token"})
			return
		}

		// Inject claims into context for handlers
		ctx := context.WithValue(r.Context(), UserClaimsKey, claims)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
