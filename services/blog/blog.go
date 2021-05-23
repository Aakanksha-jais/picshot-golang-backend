package blog

import (
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/app"

	"github.com/Aakanksha-jais/picshot-golang-backend/models"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/errors"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/types"
	"github.com/Aakanksha-jais/picshot-golang-backend/stores"
)

type blog struct {
	blogStore stores.Blog
	tagStore  stores.Tag
}

func New(blogStore stores.Blog, tagStore stores.Tag) blog {
	return blog{
		blogStore: blogStore,
		tagStore:  tagStore,
	}
}

// GetAll is used to retrieve all blogs that match the filter.
func (b blog) GetAll(c *app.Context, filter *models.Blog) ([]*models.Blog, error) {
	if filter == nil {
		filter = &models.Blog{}
	}

	return b.blogStore.GetAll(c, filter)
}

// GetAllByTagName retrieves all blogs by tag name.
func (b blog) GetAllByTagName(c *app.Context, name string) ([]*models.Blog, error) {
	tag, err := b.tagStore.GetByName(c, name)
	if err != nil {
		return nil, err
	}

	return b.blogStore.GetByIDs(c, tag.BlogIDList)
}

// GetByID is used to retrieve a single blog by its id.
func (b blog) GetByID(c *app.Context, id string) (*models.Blog, error) {
	if id == "" {
		return nil, errors.MissingParam{Param: "blog_id"}
	}

	return b.blogStore.Get(c, &models.Blog{BlogID: id})
}

// Create is used to create a Blog.
// Missing params check for fields should be done on the frontend as well.
func (b blog) Create(c *app.Context, model *models.Blog) (*models.Blog, error) {
	id := model.BlogID

	model.BlogID = "" // blog_id is automatically assigned and should remain empty before creation of blog

	err := checkMissingParams(*model)
	if err != nil {
		return nil, err
	}

	model.CreatedOn = types.Date{}.Today().String()

	// todo: store images to cloud and add image urls to model

	res, err := b.blogStore.Create(c, model)
	if err != nil {
		return nil, err
	}

	_, err = b.tagStore.AddBlogID(c, id, model.Tags)
	if err != nil {
		// tag store errors are not critical, so need not be returned to the delivery layer.
		c.Logger.Errorf("Cannot add Blog ID %s to tags %v", id, model.Tags)
	}

	return res, nil
}

func checkMissingParams(model models.Blog) error {
	if model.AccountID == 0 {
		return errors.MissingParam{Param: "account_id"}
	}

	if model.Title == "" {
		return errors.MissingParam{Param: "title"}
	}

	if model.Summary == "" {
		return errors.MissingParam{Param: "summary"}
	}

	if model.Content == "" {
		return errors.MissingParam{Param: "content"}
	}
	return nil
}

// Update updates a blog based on its id.
// Parameters that are meant to be updated are populated, else left empty.
// Images can only be added, not deleted.
func (b blog) Update(c *app.Context, model *models.Blog) (*models.Blog, error) {
	id := model.BlogID

	if id == "" {
		return nil, errors.MissingParam{Param: "blog_id"}
	}

	blog, err := b.blogStore.Get(c, &models.Blog{BlogID: id})
	if err != nil {
		return nil, errors.EntityNotFound{Entity: "blog", ID: id}
	}

	// todo: store images to cloud and add image urls to model

	res, err := b.blogStore.Update(c, model)
	if err != nil {
		return nil, err
	}

	// todo: logic to determine which tags to be removed and which to be added

	_, err = b.tagStore.RemoveBlogID(c, model.BlogID, blog.Tags)
	if err != nil {
		// tag store errors are not critical, so need not be returned to the delivery layer.
		c.Logger.Errorf("Cannot remove Blog ID %s from tags %v", id, model.Tags)
	}

	_, err = b.tagStore.AddBlogID(c, model.BlogID, model.Tags)
	if err != nil {
		// tag store errors are not critical, so need not be returned to the delivery layer.
		c.Logger.Errorf("Cannot add Blog ID %s to tags %v", id, model.Tags)
	}

	return res, nil
}

// Delete deletes a blog based on its id.
func (b blog) Delete(c *app.Context, id string) error {
	if id == "" {
		return errors.MissingParam{Param: "blog_id"}
	}

	blog, err := b.blogStore.Get(c, &models.Blog{BlogID: id})
	if err != nil {
		return errors.EntityNotFound{Entity: "blog", ID: id}
	}

	err = b.blogStore.Delete(c, id)
	if err != nil {
		return err
	}

	_, err = b.tagStore.RemoveBlogID(c, id, blog.Tags)
	if err != nil {
		// tag store errors are not critical, so need not be returned to the delivery layer.
		c.Logger.Errorf("Cannot remove Blog ID %s from tags %v", id, blog.Tags)
	}

	return nil
}
