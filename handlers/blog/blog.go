package blog

import (
	"github.com/Aakanksha-jais/picshot-golang-backend/handlers"
	"github.com/Aakanksha-jais/picshot-golang-backend/models"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/log"
	"github.com/Aakanksha-jais/picshot-golang-backend/services"
	"net/http"
)

type blog struct {
	service services.Blog
	logger  log.Logger
}

func New(service services.Blog, logger log.Logger) handlers.Blog {
	return blog{
		service: service,
		logger:  logger,
	}
}

func (b blog) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi!"))
	b.service.GetAll(r.Context(), models.Blog{})
}
