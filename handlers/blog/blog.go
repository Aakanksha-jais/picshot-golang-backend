package blog

import (
	"github.com/Aakanksha-jais/picshot-golang-backend/handlers"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/app"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/log"
	"github.com/Aakanksha-jais/picshot-golang-backend/services"
)

type blog struct {
	service services.Blog
	logger  log.Logger
}

func (b blog) GetAll(c *app.Context) (interface{}, error) {
	//todo
	return nil, nil
}

func New(service services.Blog) handlers.Blog {
	return blog{
		service: service,
	}
}
