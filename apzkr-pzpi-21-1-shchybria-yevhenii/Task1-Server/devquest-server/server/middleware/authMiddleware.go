package middleware

import (
	"devquest-server/devquest/infrastructure"
	"net/http"
	"slices"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func RolesRequired(auth infrastructure.Auth, roles string) func (http.Handler) http.Handler {
	return func (h http.Handler) http.Handler {
		return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
			token, err := auth.GetTokenFromHeader(w, r)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			requiredRoles := strings.Split(roles, ",")

			claims, _ := token.Claims.(jwt.MapClaims)
			userRole := claims["role"].(string)

			if !slices.Contains(requiredRoles, userRole) {
				w.WriteHeader(http.StatusUnauthorized)
				return
			} else {
				h.ServeHTTP(w, r)
			}
		})
	}
}