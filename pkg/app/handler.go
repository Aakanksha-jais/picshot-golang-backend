package app

import (
	"net/http"

	response2 "github.com/Aakanksha-jais/picshot-golang-backend/pkg/response"
)

type Handler func(c *Context) (interface{}, error)

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx, _ := r.Context().Value(appContextKey).(*Context)
	data, err := h(ctx)

	response2.WriteResponse(w, err, data, ctx.Logger)
}
