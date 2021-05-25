package app

import (
	"context"
)

type Context struct {
	context.Context
	Request *Request
	*App
}

func NewContext(r *Request, app *App) *Context {
	return &Context{Request: r, App: app}
}

func (c *Context) reset(r *Request) {
	c.Context = nil
	c.Request = r
}
