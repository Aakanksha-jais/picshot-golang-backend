package handlers

import (
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/app"
)

type Account interface {
	Login(c *app.Context) (interface{}, error)
	Signup(c *app.Context) (interface{}, error)
	Logout(c *app.Context) (interface{}, error)
	Get(c *app.Context) (interface{}, error)
	GetUser(c *app.Context) (interface{}, error)
	Update(c *app.Context) (interface{}, error)
	CheckAvailability(c *app.Context) (interface{}, error)
	UpdatePassword(c *app.Context) (interface{}, error)
}

type Blog interface {
	GetAll(c *app.Context) (interface{}, error)
}
