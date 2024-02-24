package middleware

import (
	"context"
	"gotemplate/internal/auth/jwt"
	"gotemplate/internal/config"
	"net/http"
	"strings"
)

const partWithToken = 1

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearerToken := r.Header.Get("Authorization")
		token := strings.Split(bearerToken, " ")[partWithToken]

		claims, err := jwt.Verify(config.Get().App.JwtSecret, token)
		if err != nil {
			// do something
		}

		userId, err := claims.GetSubject()
		if err != nil {
			// do something
		}

		ctx := context.WithValue(r.Context(), "userId", userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
