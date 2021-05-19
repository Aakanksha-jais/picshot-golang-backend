package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/auth"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/configs"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/errors"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/log"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/response"
)

func Authentication(config configs.ConfigLoader, logger log.Logger) func(inner http.Handler) http.Handler {
	return func(inner http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Debugf("request on endpoint: %s", r.URL.String())
			w.Header().Set("Content-Type", "application/json; charset=utf-8")

			if exemptPath(r) {
				inner.ServeHTTP(w, r)
				return
			}

			authHeader := strings.Split(r.Header.Get("Authorization"), " ")
			if len(authHeader) != 2 {
				logger.Errorf("invalid auth-header: %v", authHeader)
				response.New(errors.AuthError{Message: "cannot fetch auth-token; invalid auth-header"}, nil).Write(w)

				return
			}

			jwtToken := authHeader[1]

			claim, err := auth.ParseToken(config, jwtToken)
			if err != nil {
				logger.Error(err)
				response.New(err, nil).WriteHeader(w)

				return
			}

			jwtIDKey := auth.JWTContextKey("user_id")
			r = r.WithContext(context.WithValue(r.Context(), jwtIDKey, claim.UserID))
			logger.Debugf("user_id: %v authorised to make request on %v.", r.Context().Value(jwtIDKey), r.URL.Path)

			inner.ServeHTTP(w, r)
		})
	}
}

func exemptPath(req *http.Request) bool {
	url := req.URL.Path
	return strings.HasSuffix(url, "/login") || strings.HasSuffix(url, "/signup") || strings.HasSuffix(url, "/available")
}
