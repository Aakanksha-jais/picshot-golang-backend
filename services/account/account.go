package account

import (
	"context"

	"github.com/Aakanksha-jais/picshot-golang-backend/models"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/errors"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/log"
	"github.com/Aakanksha-jais/picshot-golang-backend/services"
	"github.com/Aakanksha-jais/picshot-golang-backend/stores"
	"golang.org/x/crypto/bcrypt"
)

type account struct {
	accountStore stores.Account
	blogService  services.Blog
	logger       log.Logger
}

func New(accountStore stores.Account, blogService services.Blog, logger log.Logger) services.Account {
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

func (a account) GetByID(ctx context.Context, id int64) (*models.Account, error) {
	account, err := a.accountStore.Get(ctx, &models.Account{User: models.User{ID: id}})
	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, errors.EntityNotFound{Entity: "user"}
	}

	account.Password = ""

	return account, nil
}

// GetAccountWithBlogs fetches an account with all the blogs posted by the account.
func (a account) GetAccountWithBlogs(ctx context.Context, username string) (*models.Account, error) {
	err := validateUsername(username)
	if err != nil {
		return nil, err
	}

	account, err := a.accountStore.Get(ctx, &models.Account{User: models.User{UserName: username}})
	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, errors.EntityNotFound{Entity: "user"}
	}

	blogs, err := a.blogService.GetAll(ctx, models.Blog{AccountID: account.ID})

	for i := range blogs {
		if blogs[i] != nil {
			account.Blogs = append(account.Blogs, *blogs[i])
		}
	}

	return account, nil
}

// Update updates account information based on account_id.
func (a account) Update(ctx context.Context, model *models.Account) (*models.Account, error) {
	if model.ID == 0 {
		return nil, errors.MissingParam{Param: "user_id"}
	}

	return a.accountStore.Update(ctx, model)
}

// Delete deactivates an account and updates it's deletion request.
// After 30 days, the account gets deleted if the status remains inactive.
func (a account) Delete(ctx context.Context, id int64) error {
	return a.accountStore.Delete(ctx, id)
}

// Create creates an account and assigns an id to it.
func (a account) Create(ctx context.Context, user *models.User) (*models.Account, error) {
	if user == nil {
		return nil, errors.MissingParam{Param: "user details"}
	}

	// check that the user does not exist already
	err := a.checkUserExists(ctx, user)
	if err != nil {
		return nil, err
	}

	// check if user details are valid
	err = validateDetails(user)
	if err != nil {
		return nil, err
	}

	account := models.Account{User: *user, Status: "ACTIVE"}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.Error{Err: err, Message: "error in hashing password", Type: "password-hash-error"}
	}

	account.Password = string(password)

	return a.accountStore.Create(ctx, &account)
}

// CheckAvailability checks if user name exists in the database.
func (a account) CheckAvailability(ctx context.Context, user *models.User) error {
	if empty(user) {
		return errors.MissingParam{Param: "signup_id"}
	}

	if user.UserName == "" {
		if user.Email.String == "" {
			if err := validatePhone(user.PhoneNo.String); err != nil {
				return err
			}

			return a.checkPhoneAvailability(ctx, user.PhoneNo.String)
		}

		if err := validateEmail(user.Email.String); err != nil {
			return err
		}

		return a.checkEmailAvailability(ctx, user.Email.String)
	}

	if err := validateUsername(user.UserName); err != nil {
		return err
	}

	return a.checkUsernameAvailability(ctx, user.UserName)
}

// Login gets an account by the User Details filter.
func (a account) Login(ctx context.Context, user *models.User) (*models.Account, error) {
	if user == nil {
		return nil, errors.MissingParam{Param: "user details"}
	}

	if empty(user) {
		return nil, errors.MissingParam{Param: "login_id"}
	}

	if user.UserName != "" {
		err := validateUsername(user.UserName)
		if err != nil {
			return nil, err
		}
	}

	if user.Email.String != "" {
		if err := validateEmail(user.Email.String); err != nil {
			return nil, err
		}
	}

	if user.PhoneNo.String != "" {
		if err := validatePhone(user.PhoneNo.String); err != nil {
			return nil, err
		}
	}

	account, err := a.accountStore.Get(ctx, &models.Account{User: *user})
	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, errors.EntityNotFound{Entity: "user"}
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(user.Password))
	if err != nil {
		return nil, errors.AuthError{Err: err, Message: "invalid password"}
	}

	return account, nil
}
