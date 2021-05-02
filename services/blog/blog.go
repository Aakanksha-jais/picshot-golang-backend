package blog

import (
	"context"
	"github.com/Aakanksha-jais/picshot-golang-backend/models"
	errors2 "github.com/Aakanksha-jais/picshot-golang-backend/pkg/errors"
	log2 "github.com/Aakanksha-jais/picshot-golang-backend/pkg/log"
	types2 "github.com/Aakanksha-jais/picshot-golang-backend/pkg/types"
	"github.com/Aakanksha-jais/picshot-golang-backend/services"
	"github.com/Aakanksha-jais/picshot-golang-backend/stores"
)

type blog struct {
	blogStore stores.Blog
	tagStore  stores.Tag
	logger    log2.Logger
}

func New(blogStore stores.Blog, tagStore stores.Tag, logger log2.Logger) services.Blog {
	return blog{
		blogStore: blogStore,
		tagStore:  tagStore,
		logger:    logger,
	}
}

// GetAll is used to retrieve all blogs that match the filter.
func (b blog) GetAll(ctx context.Context, filter models.Blog) ([]*models.Blog, error) {
	return b.blogStore.GetAll(ctx, filter)
}

// GetAllByTagName retrieves all blogs by tag name.
func (b blog) GetAllByTagName(ctx context.Context, name string) ([]*models.Blog, error) {
	tag, err := b.tagStore.GetByName(ctx, name)
	if err != nil {
		return nil, err
	}

	return b.blogStore.GetByIDs(ctx, tag.BlogIDList)
}

// GetByID is used to retrieve a single blog by its id.
func (b blog) GetByID(ctx context.Context, id string) (*models.Blog, error) {
	if id == "" {
		return nil, errors2.MissingParam{Param: "blog_id"}
	}

	return b.blogStore.Get(ctx, models.Blog{BlogID: id})
}

// Create is used to create a Blog.
// Missing params check for fields should be done on the frontend as well.
func (b blog) Create(ctx context.Context, model models.Blog) (*models.Blog, error) {
	id := model.BlogID

	model.BlogID = "" // blog_id is automatically assigned and should remain empty before creation of blog

	err := checkMissingParams(model)
	if err != nil {
		return nil, err
	}

	model.CreatedOn = types2.Date{}.Today().String()

	// todo: store images to cloud and add image urls to model

	res, err := b.blogStore.Create(ctx, model)
	if err != nil {
		return nil, err
	}

	_, err = b.tagStore.AddBlogID(ctx, id, model.Tags)
	if err != nil {
		// tag store errors are not critical, so need not be returned to the delivery layer.
		b.logger.Errorf("Cannot add Blog ID %s to tags %v", id, model.Tags)
	}

	return res, nil
}

func checkMissingParams(model models.Blog) error {
	if model.AccountID == 0 {
		return errors2.MissingParam{Param: "account_id"}
	}

	if model.Title == "" {
		return errors2.MissingParam{Param: "title"}
	}

	if model.Summary == "" {
		return errors2.MissingParam{Param: "summary"}
	}

	if model.Content == "" {
		return errors2.MissingParam{Param: "content"}
	}
	return nil
}

// Update updates a blog based on its id.
// Parameters that are meant to be updated are populated, else left empty.
// Images can only be added, not deleted.
func (b blog) Update(ctx context.Context, model models.Blog) (*models.Blog, error) {
	id := model.BlogID

	if id == "" {
		return nil, errors2.MissingParam{Param: "blog_id"}
	}

	blog, err := b.blogStore.Get(ctx, models.Blog{BlogID: id})
	if err != nil {
		return nil, errors2.EntityNotFound{Entity: "blog", ID: id}
	}

	// todo: store images to cloud and add image urls to model

	res, err := b.blogStore.Update(ctx, model)
	if err != nil {
		return nil, err
	}

	// todo: logic to determine which tags to be removed and which to be added

	_, err = b.tagStore.RemoveBlogID(ctx, model.BlogID, blog.Tags)
	if err != nil {
		// tag store errors are not critical, so need not be returned to the delivery layer.
		b.logger.Errorf("Cannot remove Blog ID %s from tags %v", id, model.Tags)
	}

	_, err = b.tagStore.AddBlogID(ctx, model.BlogID, model.Tags)
	if err != nil {
		// tag store errors are not critical, so need not be returned to the delivery layer.
		b.logger.Errorf("Cannot add Blog ID %s to tags %v", id, model.Tags)
	}

	return res, nil
}

// Delete deletes a blog based on its id.
func (b blog) Delete(ctx context.Context, id string) error {
	if id == "" {
		return errors2.MissingParam{Param: "blog_id"}
	}

	blog, err := b.blogStore.Get(ctx, models.Blog{BlogID: id})
	if err != nil {
		return errors2.EntityNotFound{Entity: "blog", ID: id}
	}

	err = b.blogStore.Delete(ctx, id)
	if err != nil {
		return err
	}

	_, err = b.tagStore.RemoveBlogID(ctx, id, blog.Tags)
	if err != nil {
		// tag store errors are not critical, so need not be returned to the delivery layer.
		b.logger.Errorf("Cannot remove Blog ID %s from tags %v", id, blog.Tags)
	}

	return nil
}
