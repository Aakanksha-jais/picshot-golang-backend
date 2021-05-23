package blog

import (
	"github.com/Aakanksha-jais/picshot-golang-backend/handlers"
	"github.com/Aakanksha-jais/picshot-golang-backend/models"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/app"
	"github.com/Aakanksha-jais/picshot-golang-backend/services"
)

type blog struct {
	service services.Blog
}

func (b blog) GetAll(c *app.Context) (interface{}, error) {
	return b.service.GetAll(c, models.Blog{})
}

func New(service services.Blog) handlers.Blog {
	return blog{
		service: service,
	}
}
