package services

import (
	"context"
	"github.com/Aakanksha-jais/picshot-golang-backend/models"
)

type Account interface {
	// GetAll gets all accounts that match the filter.
	GetAll(ctx context.Context, filter *models.Account) ([]*models.Account, error)

	// GetByID fetches an account with all the blogs posted by the account.
	GetByID(ctx context.Context, filter *models.Account) (*models.Account, error)

	// Create creates an account and assigns an account_id to it.
	Create(ctx context.Context, user *models.User) (*models.Account, error)

	// Update updates account information based on account_id.
	Update(ctx context.Context, model *models.Account) (*models.Account, error)

	// Delete deactivates an account and updates it's deletion request.
	// After 30 days, the account gets deleted if the status remains inactive.
	Delete(ctx context.Context, id int64) error

	// Get gets an account by the User Details filter.
	Get(ctx context.Context, user *models.User) (*models.Account, error)

	// CheckAvailability checks if username, phone number and email exist in the database already.
	CheckAvailability(ctx context.Context, user models.User) error
}

type Blog interface {
	// GetAll is used to retrieve all blogs that match the filter.
	GetAll(ctx context.Context, filter models.Blog) ([]*models.Blog, error)

	// GetAllByTagName retrieves all blogs by tag name.
	GetAllByTagName(ctx context.Context, name string) ([]*models.Blog, error)

	// GetByID is used to retrieve a single blog by its id.
	GetByID(ctx context.Context, id string) (*models.Blog, error)

	// Create is used to create a Blog.
	Create(ctx context.Context, model models.Blog) (*models.Blog, error)

	// Update updates a blog based on its id.
	// Parameters that are meant to be updated are populated, else left empty.
	// Images can only be added, not deleted.
	Update(ctx context.Context, model models.Blog) (*models.Blog, error)

	// Delete deletes a blog based on its id.
	Delete(ctx context.Context, id string) error
}
