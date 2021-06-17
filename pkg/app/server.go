package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/Aakanksha-jais/picshot-golang-backend/middlewares"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/auth"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/configs"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/log"
)

type server struct {
	contextPool sync.Pool
	Router      *Router
}

func NewServer(app *App) *server {
	s := &server{
		Router: NewRouter(),
	}

	s.setUpAuth(app.Config, app.Logger)

	s.Router.Use(s.contextInjector())

	s.contextPool.New = func() interface{} {
		return NewContext(nil, nil, app)
	}

	return s
}

func (s *server) Start(logger log.Logger) {
	server := &http.Server{
		Handler: s.Router,
		Addr:    fmt.Sprintf("localhost:%s", os.Getenv("HTTP_PORT")),
	}

	logger.Infof("starting server at PORT: %v", os.Getenv("HTTP_PORT"))

	logger.Fatalf("error in starting the server: %v", server.ListenAndServe())
}

func (s *server) setUpAuth(config configs.Config, logger log.Logger) {
	if options, ok := getOAuthOptions(config); ok {
		s.Router.Use(middlewares.Authentication(logger, options))

		logger.Infof("auth middleware enabled")
	}
}

func getOAuthOptions(config configs.Config) (auth.Options, bool) {
	var (
		options auth.Options
		ok      bool
	)

	if jwkPath := config.Get("JWKS_ENDPOINT"); jwkPath != "" {
		options.JWKPath = jwkPath
		ok = true

		if validFrequency, err := strconv.Atoi(config.Get("OAUTH_CACHE_VALIDITY")); err != nil || validFrequency == 0 {
			options.ValidityFrequency = 1800
		} else {
			options.ValidityFrequency = validFrequency
		}
	}

	return options, ok
}

type contextKey int

const appContextKey contextKey = 1

func (s *server) contextInjector() func(inner http.Handler) http.Handler {
	return func(inner http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c := s.contextPool.Get().(*Context)
			c.reset(NewRequest(r), w)
			c.Context = r.Context()

			appContext := context.WithValue(c, appContextKey, c)
			inner.ServeHTTP(w, r.WithContext(appContext))

			s.contextPool.Put(c)
		})
	}
}
