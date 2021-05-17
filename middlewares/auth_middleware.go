package middlewares

import (
	"context"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/auth"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/configs"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/errors"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/log"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/response"
	"net/http"
	"strings"
)

type JWTContextKey string

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
				response.WriteResponse(w, errors.AuthError{Message: "cannot fetch auth-token; invalid auth-header"}, logger)

				return
			}

			jwtToken := authHeader[1]

			claim, err := auth.ParseToken(config, jwtToken)
			if err != nil {
				logger.Error(err)
				response.SetHeader(w, err, logger)

				return
			}

			jwtIDKey := JWTContextKey("user_id")
			r = r.WithContext(context.WithValue(r.Context(), jwtIDKey, claim.UserID))
			logger.Infof("user_id: %v authorised.", r.Context().Value(jwtIDKey))

			inner.ServeHTTP(w, r)
		})
	}
}

func exemptPath(req *http.Request) bool {
	url := req.URL.Path
	return strings.HasSuffix(url, "/login") || strings.HasSuffix(url, "/signup") || strings.HasSuffix(url, "/available")
}
