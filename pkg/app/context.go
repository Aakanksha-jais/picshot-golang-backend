package app

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/auth"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/errors"

	err "errors"
)

type Context struct {
	context.Context
	Request *Request
	w       http.ResponseWriter
	auth.Claims
	*App
}

func NewContext(r *Request, w http.ResponseWriter, app *App) *Context {
	return &Context{Request: r, w: w, App: app, Claims: auth.NewEmptyClaim()}
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

func (c *Context) CheckAuthHeader() error {
	authHeader := c.Request.Header("Authorization")
	if authHeader == "" || len(strings.Split(authHeader, " ")) < 2 {
		return nil
	}

	jwtToken := strings.Split(authHeader, " ")[1]
	e := c.ParseToken(jwtToken)
	if e != nil {
		return nil
	}

	return errors.Error{Type: "login-error", Message: "logout before logging in"}
}
