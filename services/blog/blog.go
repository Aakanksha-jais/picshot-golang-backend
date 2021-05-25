package blog

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/auth"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/app"
	"github.com/google/uuid"

	"github.com/Aakanksha-jais/picshot-golang-backend/models"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/errors"
	"github.com/Aakanksha-jais/picshot-golang-backend/stores"
)

type blog struct {
	blogStore  stores.Blog
	tagStore   stores.Tag
	imageStore stores.Image
}

func New(blogStore stores.Blog, tagStore stores.Tag, imageStore stores.Image) blog {
	return blog{
		blogStore:  blogStore,
		tagStore:   tagStore,
		imageStore: imageStore,
	}
}

// GetAll is used to retrieve all blogs that match the filter.
func (b blog) GetAll(ctx *app.Context, filter *models.Blog) ([]*models.Blog, error) {
	if filter == nil {
		filter = &models.Blog{}
	}

	return b.blogStore.GetAll(ctx, filter)
}

// GetAllByTagName retrieves all blogs by tag name.
func (b blog) GetAllByTagName(ctx *app.Context, name string) ([]*models.Blog, error) {
	tag, err := b.tagStore.GetByName(ctx, name)
	if err != nil {
		return nil, err
	}

	return b.blogStore.GetByIDs(ctx, tag.BlogIDList)
}

// GetByID is used to retrieve a single blog by its id.
func (b blog) GetByID(ctx *app.Context, id string) (*models.Blog, error) {
	if id == "" {
		return nil, errors.MissingParam{Param: "blog_id"}
	}

	blog, err := b.blogStore.Get(ctx, &models.Blog{BlogID: id})
	if blog == nil {
		return nil, errors.EntityNotFound{Entity: "blog", ID: id}
	}

	return blog, err
}

// Create is used to create a Blog.
// Missing params check for fields should be done on the frontend as well.
func (b blog) Create(ctx *app.Context, model *models.Blog, images []*multipart.FileHeader) (*models.Blog, error) {
	id := ctx.Value(auth.JWTContextKey("user_id"))
	model.AccountID = id.(int64)

	model.BlogID = generateNewID()

	err := checkMissingParams(*model)
	if err != nil {
		return nil, err
	}

	model.CreatedOn = time.Now()

	ctx.Logger.Debugf("images to be uploaded: %v", len(images))

	for _, img := range images { //todo concurrently
		name := fmt.Sprintf("%v_%v%v", model.AccountID, generateNewID(), filepath.Ext(img.Filename))

		err := b.imageStore.Upload(ctx, img, name)
		if err != nil {
			return nil, err
		}

		model.Images = append(model.Images, fmt.Sprintf("https://%v.s3.ap-south-1.amazonaws.com/%s", ctx.Config.Get("AWS_BUCKET"), name))
	}

	res, err := b.blogStore.Create(ctx, model)
	if err != nil {
		return nil, err
	}

	_, err = b.tagStore.AddBlogID(ctx, res.BlogID, model.Tags)
	if err != nil {
		// tag store errors are not critical, so need not be returned to the delivery layer.
		ctx.Logger.Errorf("cannot add blog id %s to tags %v", res.BlogID, model.Tags)
	}

	return res, nil
}

func generateNewID() string {
	var space uuid.UUID

	space, err := uuid.NewDCEGroup()
	if err != nil {
		space, _ = uuid.NewDCEPerson()
	}

	return uuid.NewMD5(space, nil).String()
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
func (b blog) Update(ctx *app.Context, model *models.Blog) (*models.Blog, error) {
	id := model.BlogID

	if id == "" {
		return nil, errors.MissingParam{Param: "blog_id"}
	}

	blog, err := b.blogStore.Get(ctx, &models.Blog{BlogID: id})
	if err != nil {
		return nil, errors.EntityNotFound{Entity: "blog", ID: id}
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
		ctx.Logger.Errorf("Cannot remove Blog ID %s from tags %v", id, model.Tags)
	}

	_, err = b.tagStore.AddBlogID(ctx, model.BlogID, model.Tags)
	if err != nil {
		// tag store errors are not critical, so need not be returned to the delivery layer.
		ctx.Logger.Errorf("Cannot add Blog ID %s to tags %v", id, model.Tags)
	}

	return res, nil
}

// Delete deletes a blog based on its id.
func (b blog) Delete(ctx *app.Context, id string) error {
	if id == "" {
		return errors.MissingParam{Param: "blog_id"}
	}

	blog, err := b.blogStore.Get(ctx, &models.Blog{BlogID: id})
	if err != nil {
		return errors.EntityNotFound{Entity: "blog", ID: id}
	}

	err = b.blogStore.Delete(ctx, id)
	if err != nil {
		return err
	}

	_, err = b.tagStore.RemoveBlogID(ctx, id, blog.Tags)
	if err != nil {
		// tag store errors are not critical, so need not be returned to the delivery layer.
		ctx.Logger.Errorf("Cannot remove Blog ID %s from tags %v", id, blog.Tags)
	}

	return nil
}
