package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/auth"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/log"
)

const headerLength = 2

func Authentication(logger log.Logger, claims auth.Claims) func(inner http.Handler) http.Handler {
	return func(inner http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Debugf("request on endpoint: %s", r.URL.String())

			if exemptPath(r) {
				inner.ServeHTTP(w, r)
				return
			}

			authHeader := strings.Split(r.Header.Get("Authorization"), " ")
			if len(authHeader) != headerLength {
				logger.Errorf("cannot fetch auth-token (invalid auth-header): %v", authHeader)
				w.WriteHeader(http.StatusUnauthorized)

				return
			}

			jwtToken := authHeader[1]

			err := claims.ParseToken(jwtToken)
			if err != nil {
				logger.Error(err)
				w.WriteHeader(http.StatusUnauthorized)

				return
			}

			jwtIDKey := auth.JWTContextKey("user_id")
			r = r.WithContext(context.WithValue(r.Context(), jwtIDKey, claims.GetUserID()))
			logger.Debugf("user_id: %v authorized to make request on %v.", r.Context().Value(jwtIDKey), r.URL.Path)

			inner.ServeHTTP(w, r)
		})
	}
}

func exemptPath(req *http.Request) bool {
	url := req.URL.Path
	return strings.HasSuffix(url, "/login") || strings.HasSuffix(url, "/signup") || strings.HasSuffix(url, "/available")
}
