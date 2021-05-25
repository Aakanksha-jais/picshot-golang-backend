package blog

import (
	"strconv"
	"strings"

	"github.com/Aakanksha-jais/picshot-golang-backend/handlers"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/errors"

	"github.com/Aakanksha-jais/picshot-golang-backend/models"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/app"
	"github.com/Aakanksha-jais/picshot-golang-backend/services"
)

type blog struct {
	service services.Blog
}

func (b blog) Browse(ctx *app.Context) (interface{}, error) {
	panic("implement me")
}

func (b blog) Delete(ctx *app.Context) (interface{}, error) {
	id := ctx.Request.PathParam("blogid")

	return nil, b.service.Delete(ctx, id)
}

func New(service services.Blog) handlers.Blog {
	return blog{
		service: service,
	}
}

func (b blog) GetAll(ctx *app.Context) (interface{}, error) {
	return b.service.GetAll(ctx, nil)
}

func (b blog) GetBlogsByUser(ctx *app.Context) (interface{}, error) {
	accountID := ctx.Request.PathParam("accountid")

	if accountID == "" {
		return nil, errors.MissingParam{Param: "account ID"}
	}

	id, err := strconv.Atoi(accountID)
	if err != nil {
		return nil, errors.InvalidParam{Param: "account ID"}
	}

	return b.service.GetAll(ctx, &models.Blog{AccountID: int64(id)})
}

func (b blog) Get(ctx *app.Context) (interface{}, error) {
	blogID := ctx.Request.PathParam("blogid")

	return b.service.GetByID(ctx, blogID)
}

func (b blog) Create(ctx *app.Context) (interface{}, error) {
	fileHeaders := ctx.Request.ParseImages()

	blog := &models.Blog{
		Title:   ctx.Request.FormValue("title"),
		Summary: ctx.Request.FormValue("summary"),
		Content: ctx.Request.FormValue("content"),
		Tags:    strings.Split(ctx.Request.FormValue("tags"), ","),
	}

	return b.service.Create(ctx, blog, fileHeaders)
}
