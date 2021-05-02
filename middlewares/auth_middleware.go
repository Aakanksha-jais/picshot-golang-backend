package middlewares

import (
	"context"
	"github.com/Aakanksha-jais/picshot-golang-backend/handlers"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/auth"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/configs"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/errors"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/log"
	"net/http"
	"strings"
)

type JWTContextKey string

func Authentication(config configs.ConfigLoader, logger log.Logger) func(inner http.Handler) http.Handler {
	return func(inner http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Debugf("request on endpoint: %s", r.URL.String())

			if exemptPath(r) {
				inner.ServeHTTP(w, r)
				return
			}

			cookie, err := r.Cookie("auth-token")
			if err != nil || cookie == nil {
				logger.Error("cannot fetch auth-token; cookie missing")
				handlers.WriteError(w, errors.AuthError{Message: "cannot fetch auth-token; cookie missing"})
				return
			}

			jwtToken := cookie.Value

			claim, err := auth.ParseToken(config, jwtToken)
			if err != nil {
				logger.Error(err)
				handlers.SetHeader(w, err)

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
	return strings.HasSuffix(url, "/login") || strings.HasSuffix(url, "/signup")
}
