package stores

import (
	"mime/multipart"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/constants"

	"github.com/Aakanksha-jais/picshot-golang-backend/models"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/app"
)

type Account interface {
	// GetAll retrieves all accounts that match the given filter.
	GetAll(c *app.Context, filter *models.Account) ([]*models.Account, error)

	// Get retrieves a single account that matches a given filter.
	Get(c *app.Context, filter *models.Account) (*models.Account, error)

	// Create creates an account.
	Create(c *app.Context, model *models.Account) (*models.Account, error)

	// Update updates an account.
	Update(c *app.Context, model *models.Account) (*models.Account, error)

	// Delete updates a delete request for an account and sets its status to inactive.
	// Account is then permanently deleted after 30 days of inactivity.
	Delete(c *app.Context, id int64) error
}

type Blog interface {
	// GetAll is used to retrieve all blogs that match the filter.
	// BLogs can be filtered by account_id, blog_id and title.
	GetAll(c *app.Context, filter *models.Blog, page *models.Page) ([]*models.Blog, error)

	// GetByIDs retrieves all blogs whose IDs have been provided as parameter.
	GetByIDs(c *app.Context, idList []string) ([]*models.Blog, error)

	// Get is used to retrieve a SINGLE blog that matches the filter.
	// A blog can be filtered by account_id, blog_id and title.
	Get(c *app.Context, filter *models.Blog) (*models.Blog, error)

	// Create is used to create a new blog.
	Create(c *app.Context, model *models.Blog) (*models.Blog, error)

	// Update updates the blog by its ID.
	// Images and Tags can be added and not deleted todo
	Update(c *app.Context, model *models.Blog) (*models.Blog, error)

	// Delete deletes a blog by its ID.
	Delete(c *app.Context, blogID string) error
}

type Tag interface {
	// Get retrieves a tag by its name.
	// A tag name uniquely identifies a tag entity.
	// The tag entity has tag name and list of blog_id's associated with the tag.
	Get(c *app.Context, name string) (*models.Tag, error)

	// Update adds blog_id to given list of tags.
	// Tags are created if they do not exist already.
	Update(c *app.Context, blogID string, tag string, operation constants.Operation) error

	// Create creates a tag.
	Create(c *app.Context, tag *models.Tag) error

	// Delete removes a tag by its name.
	Delete(c *app.Context, tag string) error
}

type Image interface {
	// Upload uploads a file to S3 Bucket.
	Upload(c *app.Context, fileHeader *multipart.FileHeader, name string) error

	// DeleteBulk deletes multiple files whose names are passed as parameter.
	DeleteBulk(ctx *app.Context, names []string) error
}
