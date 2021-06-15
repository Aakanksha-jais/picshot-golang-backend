package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/app/response"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/errors"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/auth"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/log"
)

func Authentication(logger log.Logger, options auth.Options) func(inner http.Handler) http.Handler {
	oAuth := auth.New(options)

	return func(inner http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Debugf("request on endpoint: %s", r.URL.String())

			if err, ok := exemptPath(oAuth, r); ok {
				if err != nil {
					response.WriteResponse(w, err, nil, logger)
					return
				}

				inner.ServeHTTP(w, r)
				return
			}

			token, err := oAuth.Validate(r)
			if err == nil {
				claims := token.Claims.(*auth.Claims)

				jwtIDKey := auth.JWTContextKey("claims")
				r = r.WithContext(context.WithValue(r.Context(), jwtIDKey, claims))

				logger.Debugf("user_id: %v authorized to make request on %v.", r.Context().Value(jwtIDKey).(*auth.Claims).UserID, r.URL.Path)

				inner.ServeHTTP(w, r)
				return
			}

			response.WriteResponse(w, err, nil, logger)
		})
	}
}

func exemptPath(oAuth *auth.OAuth, req *http.Request) (error, bool) {
	url := req.URL.Path

	if strings.HasSuffix(url, "/login") || strings.HasSuffix(url, "/signup") {
		_, err := oAuth.Validate(req)
		if err != nil {
			return nil, true // login and signup allowed if auth header is invalid
		}

		return errors.Error{Type: "login-error", Msg: "logout before logging in"}, false
	}

	return nil, strings.HasSuffix(url, "/available") || strings.HasSuffix(url, "/.well-known/jwks.json")
}
