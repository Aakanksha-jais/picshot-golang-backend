package account

import (
	"context"
	"github.com/Aakanksha-jais/picshot-golang-backend/models"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/errors"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/log"
	"github.com/Aakanksha-jais/picshot-golang-backend/services"
	"github.com/Aakanksha-jais/picshot-golang-backend/stores"
	"golang.org/x/crypto/bcrypt"
	"regexp"
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
func (a account) CheckAvailability(ctx context.Context, user models.User) error {
	if user.UserName == "" {
		if user.Email.String == "" {
			if user.PhoneNo.String == "" {
				return errors.MissingParam{Param: "username (or) email (or) phone"}
			}

			if err := validatePhone(user.PhoneNo.String); err != nil {
				return err
			}

			acc, _ := a.accountStore.Get(ctx, &models.Account{User: models.User{PhoneNo: user.PhoneNo}})
			if acc != nil {
				return errors.EntityAlreadyExists{Entity: "user", ValueType: "phone_no", Value: user.PhoneNo.String}
			}

			a.logger.Debugf("phone number %s available", user.PhoneNo.String)
			return nil
		}

		if err := validateEmail(user.Email.String); err != nil {
			return err
		}

		acc, _ := a.accountStore.Get(ctx, &models.Account{User: models.User{Email: user.Email}})
		if acc != nil {
			return errors.EntityAlreadyExists{Entity: "user", ValueType: "email", Value: user.Email.String}
		}

		a.logger.Debugf("email %s available", user.Email.String)
		return nil
	}

	if err := validateUsername(user.UserName); err != nil {
		return err
	}

	acc, _ := a.accountStore.Get(ctx, &models.Account{User: models.User{UserName: user.UserName}})
	if acc != nil {
		return errors.EntityAlreadyExists{Entity: "user", ValueType: "username", Value: user.UserName}
	}

	a.logger.Debugf("username %s available", user.UserName)
	return nil
}

// Get gets an account by the User Details filter.
func (a account) Get(ctx context.Context, user *models.User) (*models.Account, error) {
	if user == nil {
		return nil, errors.MissingParam{Param: "user details"}
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

func (a account) checkUserExists(ctx context.Context, user *models.User) error {
	acc, _ := a.accountStore.Get(ctx, &models.Account{User: models.User{UserName: user.UserName}})
	if acc != nil {
		return errors.EntityAlreadyExists{Entity: "user", ValueType: "username", Value: user.UserName}
	}

	acc, _ = a.accountStore.Get(ctx, &models.Account{User: models.User{Email: user.Email}})
	if acc != nil {
		return errors.EntityAlreadyExists{Entity: "user", ValueType: "email", Value: user.Email.String}
	}

	acc, _ = a.accountStore.Get(ctx, &models.Account{User: models.User{PhoneNo: user.PhoneNo}})
	if acc != nil {
		return errors.EntityAlreadyExists{Entity: "user", ValueType: "phone number", Value: user.PhoneNo.String}
	}

	return nil
}

func validateDetails(user *models.User) error {
	if user.UserName == "" {
		return errors.MissingParam{Param: "user_name"}
	}

	if err := validateUsername(user.UserName); err != nil {
		return err
	}

	if user.Email.String == "" && user.PhoneNo.String == "" {
		return errors.MissingParam{Param: "email"}
	}

	if user.Email.String != "" {
		if err := validateEmail(user.Email.String); err != nil {
			return err
		}
	}

	if user.PhoneNo.String != "" {
		if err := validatePhone(user.PhoneNo.String); err != nil {
			return err
		}
	}

	if err := validatePassword(user.Password); err != nil {
		return err
	}

	if user.FName == "" && user.LName == "" {
		return errors.MissingParam{Param: "name"}
	}

	if err := validateName(user.FName); err != nil {
		return err
	}

	err := validateName(user.LName)

	return err
}

func validateName(name string) error {
	// username should be aplha-numeric and should have at least 8 characters
	res, err := regexp.MatchString("^[a-zA-Z]+$", name)
	if err == nil && res {
		return nil
	}

	return errors.InvalidParam{Param: "name"}
}

func validateUsername(username string) error {
	// username should be aplha-numeric and should have at least 8 characters
	res, err := regexp.MatchString(`^[0-9A-Za-z_]{8,}$`, username)
	if err == nil && res {
		return nil
	}

	return errors.InvalidParam{Param: "username"}
}

func validateEmail(email string) error {
	res, err := regexp.MatchString(`^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+[.][a-zA-Z0-9-.]+$`, email)
	if err == nil && res {
		return nil
	}

	return errors.InvalidParam{Param: "email"}
}

func validatePhone(phone string) error {
	res, err := regexp.MatchString(`^[0-9]+$`, phone)
	if err == nil && res {
		return nil
	}

	return errors.InvalidParam{Param: "phone_no"}
}

func validatePassword(password string) error {
	// password between 8 to 20 characters
	// alphanumeric and !@#$%^&* symbols allowed
	res, err := regexp.MatchString(`^[a-zA-Z0-9!@#$%^&*]{8,20}$`, password)
	if err == nil && res {
		return nil
	}

	return errors.InvalidParam{Param: "password"}
}
