package account

import (
	"database/sql"
	"time"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/constants"

	"github.com/Aakanksha-jais/picshot-golang-backend/models"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/auth"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/app"

	"github.com/Aakanksha-jais/picshot-golang-backend/handlers"
	"github.com/Aakanksha-jais/picshot-golang-backend/services"
)

type account struct {
	service services.Account
}

func (a account) Login(c *app.Context) (interface{}, error) {
	exp := time.Now().Add(30 * time.Minute)

	user, err := c.Request.UnmarshalUser()
	if err != nil {
		return nil, err
	}

	account, err := a.service.Login(c, user)
	if err != nil {
		return nil, err
	}

	token, err := auth.CreateToken(auth.NewClaim(exp.Unix(), account.ID))
	if err != nil {
		c.Logger.Warnf("error in generating token: %v", err)
	} else {
		c.Logger.Debugf("token generated successfully \033[32m[expires at: %v]\033[0m", exp.Format(time.RFC850))

		c.SetAuthHeader(token)
	}

	return nil, err
}

func (a account) Signup(c *app.Context) (interface{}, error) {
	exp := time.Now().Add(30 * time.Minute)

	user, err := c.Request.UnmarshalUser()
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("signup request for %v", user)

	account, err := a.service.Create(c, user)
	if err != nil {
		return nil, err
	}

	token, err := auth.CreateToken(auth.NewClaim(exp.Unix(), account.ID))
	if err != nil {
		c.Logger.Warnf("error in generating token: %v", err)
	} else {
		c.Logger.Debugf("token generated successfully \033[32m[expires at: %v]\u001B[0m", exp.Format(time.RFC850))

		c.SetAuthHeader(token)
	}

	return constants.CreateSuccess, err
}

func (a account) Logout(c *app.Context) (interface{}, error) {
	return nil, nil
}

func (a account) Get(c *app.Context) (interface{}, error) {
	id := c.Value(auth.JWTContextKey("user_id"))

	c.Logger.Debugf("user with id: %v is logged in", id)

	return a.service.GetByID(c, id.(int64))
}

func (a account) GetUser(c *app.Context) (interface{}, error) {
	username := c.Request.PathParam("username")

	return a.service.GetAccountWithBlogs(c, username)
}

// Update updates user details but not password.
func (a account) Update(c *app.Context) (interface{}, error) {
	user, err := c.Request.UnmarshalUser()
	if err != nil {
		return nil, err
	}

	return a.service.UpdateUser(c, user)
}

func (a account) UpdatePassword(c *app.Context) (interface{}, error) {
	pwd := struct {
		Old string `json:"old_password"`
		New string `json:"new_password"`
	}{}

	err := c.Request.Unmarshal(&pwd)
	if err != nil {
		return nil, err
	}

	return nil, a.service.UpdatePassword(c, pwd.Old, pwd.New)
}

func (a account) CheckAvailability(c *app.Context) (interface{}, error) {
	username := c.Request.QueryParam("username")
	email := c.Request.QueryParam("email")
	phone := c.Request.QueryParam("phone")

	return nil, a.service.CheckAvailability(c, &models.User{UserName: username, PhoneNo: sql.NullString{String: phone, Valid: true}, Email: sql.NullString{String: email, Valid: true}})
}

func New(service services.Account) handlers.Account {
	return account{service: service}
}
