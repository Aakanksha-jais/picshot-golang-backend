package app

import (
	"context"
	err "errors"
	"fmt"
	"net/http"
)

type Context struct {
	context.Context
	Request *Request
	w       http.ResponseWriter
	*App
}

func NewContext(r *Request, w http.ResponseWriter, app *App) *Context {
	return &Context{Request: r, w: w, App: app}
}

func (c *Context) reset(r *Request, w http.ResponseWriter) {
	c.Context = nil
	c.Request = r
	c.w = w
}

func (c *Context) SetAuthHeader(token string) {
	c.w.Header().Set("Content-Type", "application/json; charset=utf-8")
	c.w.Header().Add("Authorization", fmt.Sprintf("Bearer %s", token))
}

func (c *Context) SetStatus(status int) error {
	c.w.WriteHeader(status)

	return err.New("do not set header")
}
