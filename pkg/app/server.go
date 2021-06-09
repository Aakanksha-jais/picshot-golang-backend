package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/auth"

	"github.com/Aakanksha-jais/picshot-golang-backend/middlewares"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/log"
)

type server struct {
	contextPool sync.Pool
	Router      *Router
}

func (s *server) Start(logger log.Logger) {
	server := &http.Server{
		Handler: s.Router,
		Addr:    fmt.Sprintf("localhost:%s", os.Getenv("HTTP_PORT")),
	}

	logger.Infof("starting server at PORT: %v", os.Getenv("HTTP_PORT"))

	logger.Fatalf("error in starting the server: %v", server.ListenAndServe())
}

func NewServer(app *App) *server {
	s := &server{
		Router: NewRouter(),
	}

	if app.Config.GetOrDefault("ENABLE_AUTH", "YES") == "YES" {
		app.Infof("authentication middleware enabled")
		s.Router.Use(middlewares.Authentication(app.Logger, auth.NewEmptyClaim()))
	} else {
		app.Warnf("authentication middleware disabled, some endpoints will not run")
	}

	s.Router.Use(s.contextInjector())

	s.contextPool.New = func() interface{} {
		return NewContext(nil, nil, app)
	}

	return s
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
