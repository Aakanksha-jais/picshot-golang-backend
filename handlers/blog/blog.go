package blog

import (
	"strconv"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/errors"

	"github.com/Aakanksha-jais/picshot-golang-backend/models"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/app"
	"github.com/Aakanksha-jais/picshot-golang-backend/services"
)

type blog struct {
	service services.Blog
}

func New(service services.Blog) blog {
	return blog{
		service: service,
	}
}

func (b blog) GetAll(c *app.Context) (interface{}, error) {
	return b.service.GetAll(c, nil)
}

func (b blog) GetBlogsByUser(c *app.Context) (interface{}, error) {
	accountID := c.Request.PathParam("accountid")
	id, err := strconv.Atoi(accountID)
	if err != nil {
		return nil, errors.InvalidParam{Param: "account ID"}
	}

	return b.service.GetAll(c, &models.Blog{AccountID: int64(id)})
}
