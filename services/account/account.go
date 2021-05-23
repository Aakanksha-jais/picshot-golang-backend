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
func (a account) GetAll(c *app.Context, filter *models.Account) ([]*models.Account, error) {
	return a.accountStore.GetAll(c, filter)
}

func (a account) GetByID(c *app.Context, id int64) (*models.Account, error) {
	account, err := a.accountStore.Get(c, &models.Account{User: models.User{ID: id}})
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
func (a account) GetAccountWithBlogs(c *app.Context, username string) (*models.Account, error) {
	err := validateUsername(username)
	if err != nil {
		return nil, err
	}

	account, err := a.accountStore.Get(c, &models.Account{User: models.User{UserName: username}})
	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, errors.EntityNotFound{Entity: "user"}
	}

	blogs, err := a.blogService.GetAll(c, &models.Blog{AccountID: account.ID})

	for i := range blogs {
		if blogs[i] != nil {
			account.Blogs = append(account.Blogs, *blogs[i])
		}
	}

	return account, nil
}

func (a account) UpdateUser(c *app.Context, user *models.User) (*models.Account, error) {
	if user == nil {
		return nil, errors.MissingParam{Param: "user details"}
	}

	update := &models.Account{}

	id := c.Value(auth.JWTContextKey("user_id"))

	account, err := a.GetByID(c, id.(int64))
	if err != nil {
		return nil, err
	}

	if account.UserName != user.UserName {
		err = validateUsername(user.UserName)
		if err != nil {
			return nil, err
		}

		err = a.CheckAvailability(c, &models.User{UserName: user.UserName})
		if err != nil {
			return nil, err
		}

		update.UserName = account.UserName
	}

	if account.FName != user.FName {
		err = validateName(user.FName)
		if err != nil {
			return nil, err
		}

		update.FName = account.FName
	}

	if account.LName != user.LName {
		err = validateName(user.LName)
		if err != nil {
			return nil, err
		}

		update.LName = account.LName
	}

	if account.Email.String != user.Email.String {
		err = validateEmail(user.Email.String)
		if err != nil {
			return nil, err
		}

		err = a.CheckAvailability(c, &models.User{Email: user.Email})
		if err != nil {
			return nil, err
		}

		update.Email = account.Email
	}

	if account.PhoneNo.String != user.PhoneNo.String {
		err = validatePhone(user.PhoneNo.String)
		if err != nil {
			return nil, err
		}

		err = a.CheckAvailability(c, &models.User{PhoneNo: user.PhoneNo})
		if err != nil {
			return nil, err
		}

		update.PhoneNo = account.PhoneNo
	}

	update.ID = id.(int64)

	return a.accountStore.Update(c, update)
}

// Update updates account information based on account_id.todo
func (a account) Update(c *app.Context, model *models.Account) (*models.Account, error) {
	if model.ID == 0 {
		return nil, errors.MissingParam{Param: "user_id"}
	}

	return a.accountStore.Update(c, model)
}

func (a account) UpdatePassword(c *app.Context, oldPassword, newPassword string) error {
	id := c.Value(auth.JWTContextKey("user_id"))

	account, err := a.accountStore.Get(c, &models.Account{User: models.User{ID: id.(int64)}})
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

	_, err = a.accountStore.Update(c, &models.Account{User: models.User{ID: id.(int64), Password: string(hash)}, PwdUpdate: &sql.NullTime{Time: time.Now(), Valid: true}})

	return err
}

// Delete deactivates an account and updates it's deletion request.
// After 30 days, the account gets deleted if the status remains inactive.
func (a account) Delete(c *app.Context, id int64) error {
	return a.accountStore.Delete(c, id)
}

// Create creates an account and assigns an id to it.
func (a account) Create(c *app.Context, user *models.User) (*models.Account, error) {
	if user == nil {
		return nil, errors.MissingParam{Param: "user details"}
	}

	// check that the user does not exist already
	err := a.checkUserExists(c, user)
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

	return a.accountStore.Create(c, &account)
}

// CheckAvailability checks if user name exists in the database.
func (a account) CheckAvailability(c *app.Context, user *models.User) error {
	if empty(user) {
		return errors.MissingParam{Param: "signup_id"}
	}

	if user.UserName == "" {
		if user.Email.String == "" {
			if err := validatePhone(user.PhoneNo.String); err != nil {
				return err
			}

			return a.checkPhoneAvailability(c, user.PhoneNo.String)
		}

		if err := validateEmail(user.Email.String); err != nil {
			return err
		}

		return a.checkEmailAvailability(c, user.Email.String)
	}

	if err := validateUsername(user.UserName); err != nil {
		return err
	}

	return a.checkUsernameAvailability(c, user.UserName)
}

// Login gets an account by the User Details filter.
func (a account) Login(c *app.Context, user *models.User) (*models.Account, error) {
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

	account, err := a.accountStore.Get(c, &models.Account{User: *user})
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
