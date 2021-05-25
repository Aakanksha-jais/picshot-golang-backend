package services

import (
	"mime/multipart"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/app"

	"github.com/Aakanksha-jais/picshot-golang-backend/models"
)

type Account interface {
	// GetAll retrieves all accounts that match the filter.
	GetAll(c *app.Context, filter *models.Account) ([]*models.Account, error)

	// GetByID retrieves an account by its id
	GetByID(c *app.Context, id int64) (*models.Account, error)

	// GetAccountWithBlogs fetches an account with all the blogs posted by the account.
	GetAccountWithBlogs(c *app.Context, username string) (*models.Account, error)

	// Create creates an account and assigns an user_id to it.
	Create(c *app.Context, user *models.User) (*models.Account, error)

	// UpdateUser updates user information based on user_id.
	UpdateUser(c *app.Context, model *models.User) (*models.Account, error)

	UpdatePassword(c *app.Context, oldPassword, newPassword string) error

	// Update updates Account info based on user_id.
	Update(c *app.Context, model *models.Account) (*models.Account, error)

	// Delete deactivates an account and updates it's deletion request.
	// After 30 days, the account gets deleted if the status remains inactive.
	Delete(c *app.Context, id int64) error

	// Login logs in a user to his account.
	Login(c *app.Context, user *models.User) (*models.Account, error)

	// CheckAvailability checks if username, phone number and email exist in the database already.
	CheckAvailability(c *app.Context, user *models.User) error
}

type Blog interface {
	// GetAll retrieve all blogs that match the filter.
	GetAll(c *app.Context, filter *models.Blog) ([]*models.Blog, error)

	// GetAllByTagName retrieves all blogs by tag name.
	GetAllByTagName(c *app.Context, name string) ([]*models.Blog, error)

	// GetByID retrieves a single blog by its id.
	GetByID(c *app.Context, id string) (*models.Blog, error)

	// Create creates a Blog.
	Create(c *app.Context, model *models.Blog, images []*multipart.FileHeader) (*models.Blog, error)

	// Update updates a blog based on its id.
	// Parameters that are meant to be updated are populated, else left empty.
	// Images can only be added, not deleted.
	Update(c *app.Context, model *models.Blog) (*models.Blog, error)

	// Delete deletes a blog based on its id.
	Delete(c *app.Context, id string) error
}
