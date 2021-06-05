package blog

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"github.com/Aakanksha-jais/picshot-golang-backend/services"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/auth"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/app"
	"github.com/google/uuid"

	"github.com/Aakanksha-jais/picshot-golang-backend/models"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/errors"
	"github.com/Aakanksha-jais/picshot-golang-backend/stores"
)

type blog struct {
	blogStore  stores.Blog
	tagService services.Tag
	imageStore stores.Image
}

func New(blogStore stores.Blog, tagService services.Tag, imageStore stores.Image) services.Blog {
	return blog{
		blogStore:  blogStore,
		tagService: tagService,
		imageStore: imageStore,
	}
}

// GetAll is used to retrieve all blogs that match the filter.
func (b blog) GetAll(ctx *app.Context, filter *models.Blog, page *models.Page) ([]*models.Blog, error) {
	if filter == nil {
		filter = &models.Blog{}
	}

	return b.blogStore.GetAll(ctx, filter, page)
}

// GetAllByTagName retrieves all blogs by tag input.
func (b blog) GetAllByTagName(ctx *app.Context, name string) ([]*models.Blog, error) {
	tag, err := b.tagService.Get(ctx, name)
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

	err := checkMissingParams(model)
	if err != nil {
		return nil, err
	}

	model.CreatedOn = time.Now()

	ctx.Debugf("images to be uploaded: %v", len(images))

	n := len(images)
	errs := make(chan error, n)

	for _, img := range images {
		name := fmt.Sprintf("%v_%v%v", model.AccountID, generateNewID(), filepath.Ext(img.Filename))

		go func() {
			errs <- b.imageStore.Upload(ctx, img, name)
		}()

		model.Images = append(model.Images, fmt.Sprintf("https://%v.s3.ap-south-1.amazonaws.com/%s", ctx.Config.Get("AWS_BUCKET"), name))
	}

	for i := 0; i < n; i++ {
		//nolint:govet // redeclare error
		if err := <-errs; err != nil {
			return nil, err
		}
	}

	res, err := b.blogStore.Create(ctx, model)
	if err != nil {
		return nil, err
	}

	b.tagService.AddBlogID(ctx, res.BlogID, model.Tags)

	return res, nil
}

// Update updates a blog based on its id.
// Images can only be added, not deleted.
// Image array will be empty if no new images are added
// Tags will be overwritten
func (b blog) Update(ctx *app.Context, model *models.Blog, images []*multipart.FileHeader) (*models.Blog, error) {
	userID := ctx.Value(auth.JWTContextKey("user_id"))
	model.AccountID = userID.(int64)

	id := model.BlogID
	if id == "" {
		return nil, errors.MissingParam{Param: "blog_id"}
	}

	err := checkMissingParams(model)
	if err != nil {
		return nil, err
	}

	// check if the blog exists already
	blog, err := b.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	n := len(images)
	errs := make(chan error, n)

	// upload new images to s3 bucket
	for _, img := range images {
		name := fmt.Sprintf("%v_%v%v", model.AccountID, generateNewID(), filepath.Ext(img.Filename))

		go func() {
			errs <- b.imageStore.Upload(ctx, img, name)
		}()

		model.Images = append(model.Images, fmt.Sprintf("https://%v.s3.ap-south-1.amazonaws.com/%s", ctx.Config.Get("AWS_BUCKET"), name))
	}

	for i := 0; i < n; i++ {
		if err := <-errs; err != nil {
			return nil, err
		}
	}

	// update blog
	res, err := b.blogStore.Update(ctx, model)
	if err != nil {
		return nil, err
	}

	b.updateTags(ctx, model, blog)

	return res, nil
}

func (b blog) updateTags(ctx *app.Context, new, old *models.Blog) {
	var addTags, removeTags []string

	switch {
	case len(new.Tags) == 0:
		removeTags = old.Tags
	case len(old.Tags) == 0:
		addTags = new.Tags
	default:
		m := make(map[string]int)

		for _, tag := range new.Tags {
			m[tag] += 1
		}

		for _, tag := range new.Tags {
			m[tag] -= 1
		}

		for tag, index := range m {
			if index == -1 {
				removeTags = append(removeTags, tag)
			}

			if index == 1 {
				addTags = append(addTags, tag)
			}
		}
	}

	b.tagService.RemoveBlogID(ctx, old.BlogID, removeTags)

	b.tagService.AddBlogID(ctx, old.BlogID, addTags)
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

	names := getNames(blog.Images)

	err = b.imageStore.DeleteBulk(ctx, names)
	if err != nil {
		return err
	}

	b.tagService.RemoveBlogID(ctx, id, blog.Tags)

	return nil
}

func getNames(images []string) []string {
	names := make([]string, 0)

	for _, img := range images {
		s := strings.Split(img, "/")
		name := s[len(s)-1]

		names = append(names, name)
	}

	return names
}

func generateNewID() string {
	var space uuid.UUID

	space, err := uuid.NewDCEGroup()
	if err != nil {
		space, _ = uuid.NewDCEPerson()
	}

	return uuid.NewMD5(space, nil).String()
}

func checkMissingParams(model *models.Blog) error {
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
