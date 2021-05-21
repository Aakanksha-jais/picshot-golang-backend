package app

import (
	"github.com/gorilla/mux"
)

type Router struct {
	mux.Router
}

// NewRouter returns a new router instance.
func NewRouter() *Router {
	return &Router{Router: *mux.NewRouter()}
}

func (r *Router) Add(method, pattern string, handler Handler) {
	r.Router.NewRoute().Methods(method).Path(pattern).Handler(handler)
}
