package app

import (
	"context"
	"fmt"
	"net/http"
)

type Context struct {
	context.Context
	Request  *Request
	Response http.ResponseWriter
	*App
}

func NewContext(r *Request, w http.ResponseWriter, app *App) *Context {
	return &Context{Request: r, Response: w, App: app}
}

func (c *Context) SetAuthHeader(token string) {
	c.Response.Header().Add("Authorization", fmt.Sprintf("Bearer %s", token))
}

func (c *Context) reset(r *Request, w http.ResponseWriter) {
	c.Context = nil
	c.Request = r
	c.Response = w
}
