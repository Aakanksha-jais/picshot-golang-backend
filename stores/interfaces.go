package stores

import (
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
	GetAll(c *app.Context, filter *models.Blog) ([]*models.Blog, error)

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
	// GetByName retrieves a tag by its name.
	// A tag name uniquely identifies a tag entity.
	// The tag entity has tag name and list of blog_id's associated with the tag.
	GetByName(c *app.Context, name string) (*models.Tag, error)

	// AddBlogID adds blog_id to given list of tags.
	// Tags are created if they do not exist already.
	AddBlogID(c *app.Context, blogID string, tags []string) ([]*models.Tag, error)

	// RemoveBlogID removes blog_id from given list of tags.
	RemoveBlogID(c *app.Context, blogID string, tags []string) ([]*models.Tag, error)
}
