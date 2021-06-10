package account

import (
	"database/sql"
	"reflect"
	"time"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/auth"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/app"

	"github.com/Aakanksha-jais/picshot-golang-backend/models"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/errors"
	"github.com/Aakanksha-jais/picshot-golang-backend/services"
	"github.com/Aakanksha-jais/picshot-golang-backend/stores"
	"golang.org/x/crypto/bcrypt"
)

type account struct {
	accountStore stores.Account
	blogService  services.Blog
}

func New(accountStore stores.Account, blogService services.Blog) services.Account {
	return account{
		accountStore: accountStore,
		blogService:  blogService,
	}
}

// GetAll gets all accounts that match the filter.
func (a account) GetAll(ctx *app.Context, filter *models.Account) ([]*models.Account, error) {
	return a.accountStore.GetAll(ctx, filter)
}

func (a account) GetByID(ctx *app.Context, id int64) (*models.Account, error) {
	account, err := a.accountStore.Get(ctx, &models.Account{User: models.User{ID: id}})
	if err != nil {
		return nil, err
	}

	if account == nil || reflect.DeepEqual(account, &models.Account{}) {
		return nil, errors.EntityNotFound{Entity: "user"}
	}

	account.Password = ""

	return account, nil
}

// GetAccountWithBlogs fetches an account with all the blogs posted by the account.
func (a account) GetAccountWithBlogs(ctx *app.Context, username string) (*models.Account, error) {
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

	blogs, err := a.blogService.GetAll(ctx, &models.Blog{AccountID: account.ID}, nil)
	if err != nil {
		return nil, err
	}

	for i := range blogs {
		if blogs[i] != nil {
			account.Blogs = append(account.Blogs, *blogs[i])
		}
	}

	return account, nil
}

func (a account) UpdateUser(ctx *app.Context, user *models.User) (*models.Account, error) {
	if user == nil {
		return nil, errors.MissingParam{Param: "user details"}
	}

	id := ctx.Value(auth.JWTContextKey("user_id"))

	account, err := a.GetByID(ctx, id.(int64))
	if err != nil {
		return nil, err
	}

	update, err := a.getUpdate(ctx, account, user)
	if err != nil {
		return nil, err
	}

	update.ID = id.(int64)

	return a.accountStore.Update(ctx, update)
}

// Update updates account information based on account_id.todo
func (a account) Update(ctx *app.Context, model *models.Account, id int64) (*models.Account, error) {
	model.ID = id

	return a.accountStore.Update(ctx, model)
}

func (a account) UpdatePassword(ctx *app.Context, oldPassword, newPassword string) error {
	id := ctx.Value(auth.JWTContextKey("user_id"))

	account, err := a.accountStore.Get(ctx, &models.Account{User: models.User{ID: id.(int64)}})
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(oldPassword))
	if err != nil {
		return errors.AuthError{Err: err, Message: "invalid password"}
	}

	err = validatePassword(newPassword)
	if err != nil {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return errors.Error{Err: err, Message: "error in hashing password", Type: "password-hash-error"}
	}

	_, err = a.accountStore.Update(ctx, &models.Account{
		User:      models.User{ID: id.(int64), Password: string(hash)},
		PwdUpdate: sql.NullTime{Time: time.Now(), Valid: true},
	})

	return err
}

// Delete deactivates an account and updates it's deletion request.
// After 30 days, the account gets deleted if the status remains inactive.
func (a account) Delete(ctx *app.Context) error { // TODO: trigger a cronjob for 30 days deletion functionality
	id := ctx.Value(auth.JWTContextKey("user_id"))

	return a.accountStore.Delete(ctx, id.(int64))
}

// Create creates an account and assigns an id to it.
func (a account) Create(ctx *app.Context, user *models.User) (*models.Account, error) {
	if user == nil {
		return nil, errors.MissingParam{Param: "user details"}
	}

	// check that the user does not exist already
	err := a.checkUserExists(ctx, user)
	if err != nil {
		return nil, err
	}

	// check if user details are valid
	err = validateUser(user)
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
func (a account) CheckAvailability(ctx *app.Context, user *models.User) error {
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
func (a account) Login(ctx *app.Context, user *models.User) (*models.Account, error) {
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

	return a.Update(ctx, &models.Account{DelRequest: sql.NullTime{}, Status: "ACTIVE"}, account.ID)
}
