package handlers

import "github.com/Aakanksha-jais/picshot-golang-backend/pkg/app"

type Account interface {
	Login(ctx *app.Context) (interface{}, error)
	Signup(ctx *app.Context) (interface{}, error)
	Logout(ctx *app.Context) (interface{}, error)
	Get(ctx *app.Context) (interface{}, error)
	GetUser(ctx *app.Context) (interface{}, error)
	Update(ctx *app.Context) (interface{}, error)
	UpdatePassword(ctx *app.Context) (interface{}, error)
	Delete(ctx *app.Context) (interface{}, error)
	CheckAvailability(ctx *app.Context) (interface{}, error)
}

type Blog interface {
	GetAll(ctx *app.Context) (interface{}, error)
	GetBlogsByUser(ctx *app.Context) (interface{}, error)
	Get(ctx *app.Context) (interface{}, error)
	Browse(ctx *app.Context) (interface{}, error)
	Delete(ctx *app.Context) (interface{}, error)
	Create(ctx *app.Context) (interface{}, error)
	Update(ctx *app.Context) (interface{}, error)
}
