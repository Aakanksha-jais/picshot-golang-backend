package middlewares

import (
	"context"
	"github.com/Aakanksha-jais/picshot-golang-backend/handlers"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/auth"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/errors"
	"net/http"
	"strings"
)

type JWTContextKey string

func Authentication(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if exemptPath(r) {
			inner.ServeHTTP(w, r)
			return
		}

		cookie, err := r.Cookie("auth-token")
		if err != nil || cookie == nil {
			handlers.WriteError(w, errors.AuthError{Message: "cannot fetch auth-token; cookie missing"})
			return
		}

		jwtToken := cookie.Value

		claim, err := auth.ParseToken(jwtToken)
		if err != nil {
			handlers.SetHeader(w, err)

			return
		}

		jwtIDKey := JWTContextKey("id")
		r = r.WithContext(context.WithValue(r.Context(), jwtIDKey, claim.UserID))

		inner.ServeHTTP(w, r)
	})
}

func exemptPath(req *http.Request) bool {
	url := req.URL.Path
	return strings.HasSuffix(url, "/login") || strings.HasSuffix(url, "/signup")
}
