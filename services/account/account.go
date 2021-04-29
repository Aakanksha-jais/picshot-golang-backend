package account

import (
	"context"
	"github.com/Aakanksha-jais/picshot-golang-backend/log"
	"github.com/Aakanksha-jais/picshot-golang-backend/models"
	"github.com/Aakanksha-jais/picshot-golang-backend/services"
	"github.com/Aakanksha-jais/picshot-golang-backend/stores"
	"golang.org/x/crypto/bcrypt"
)

type account struct {
	accountStore stores.Account
	blogService  services.Blog
	logger       log.Logger
}

func New(accountStore stores.Account, blogService services.Blog, logger log.Logger) account {
	return account{
		accountStore: accountStore,
		blogService:  blogService,
		logger:       logger,
	}
}

// GetAll gets all accounts that match the filter.
func (a account) GetAll(ctx context.Context, filter *models.Account) ([]*models.Account, error) {
	return a.accountStore.GetAll(ctx, filter)
}

// GetByID fetches an account with all the blogs posted by the account.
func (a account) GetByID(ctx context.Context, filter *models.Account) (*models.Account, error) {
	account, err := a.accountStore.Get(ctx, filter)
	if err != nil {
		return nil, err
	}

	blogs, err := a.blogService.GetAll(ctx, models.Blog{AccountID: account.ID})

	for i := range blogs {
		if blogs[i] != nil {
			account.Blogs = append(account.Blogs, *blogs[i])
		}
	}

	return account, nil
}

// Create creates an account and assigns an account_id to it.
func (a account) Create(ctx context.Context, model *models.Account) (*models.Account, error) {
	password, err := bcrypt.GenerateFromPassword(model.Password, bcrypt.MaxCost)
	if err != nil {
		//todo return error
	}

	model.Password = password

	return a.accountStore.Create(ctx, model)
}

// Update updates account information based on account_id.
func (a account) Update(ctx context.Context, model *models.Account) (*models.Account, error) {
	return a.accountStore.Update(ctx, model)
}

// Delete deactivates an account and updates it's deletion request.
// After 30 days, the account gets deleted if the status remains inactive.
func (a account) Delete(ctx context.Context, id int64) error {
	return a.accountStore.Delete(ctx, id)
}

func (a account) LogIn(ctx context.Context, user models.User) error {
	account, err := a.GetByID(ctx, &models.Account{User: user})
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword(account.Password, user.Password)
	if err != nil {
		//todo incorrect password error
	}

	return nil
}
