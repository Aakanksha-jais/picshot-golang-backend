package account

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/Aakanksha-jais/picshot-golang-backend/handlers"

	"github.com/Aakanksha-jais/picshot-golang-backend/models"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/auth"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/app"

	"github.com/Aakanksha-jais/picshot-golang-backend/services"
)

type account struct {
	service services.Account
}

func New(service services.Account) handlers.Account {
	return account{service: service}
}

func (a account) Login(ctx *app.Context) (interface{}, error) {
	err := ctx.CheckAuthHeader()
	if err != nil {
		return nil, err
	}

	exp := time.Now().Add(24 * time.Hour)

	user, err := ctx.Request.UnmarshalUser()
	if err != nil {
		return nil, err
	}

	account, err := a.service.Login(ctx, user)
	if err != nil {
		return nil, err
	}

	ctx.Claims = auth.New(exp.Unix(), account.ID)

	token, err := ctx.Claims.CreateToken()
	if err != nil {
		return nil, err
	}

	ctx.Debugf("token generated successfully \033[32m[expires at: %v]\033[0m", exp.Format(time.RFC850))

	ctx.SetAuthHeader(token)

	return &account.User, ctx.SetStatus(http.StatusOK)
}

func (a account) Signup(ctx *app.Context) (interface{}, error) {
	err := ctx.CheckAuthHeader()
	if err != nil {
		return nil, err
	}

	exp := time.Now().Add(24 * time.Hour)

	user, err := ctx.Request.UnmarshalUser()
	if err != nil {
		return nil, err
	}

	ctx.Debugf("signup request for %v", user)

	account, err := a.service.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	ctx.Claims = auth.New(exp.Unix(), account.ID)

	token, err := ctx.Claims.CreateToken()
	if err != nil {
		return nil, err
	}

	ctx.Debugf("token generated successfully \033[32m[expires at: %v]\033[0m", exp.Format(time.RFC850))

	ctx.SetAuthHeader(token)

	return &account.User, ctx.SetStatus(http.StatusCreated)
}

func (a account) Logout(ctx *app.Context) (interface{}, error) { //todo invalidate the auth token
	ctx.SetAuthHeader("")

	return nil, ctx.SetStatus(http.StatusOK)
}

func (a account) Get(ctx *app.Context) (interface{}, error) {
	id := ctx.Value(auth.JWTContextKey("user_id"))

	ctx.Debugf("user with id: %v is logged in", id)

	return a.service.GetByID(ctx, id.(int64))
}

func (a account) GetUser(ctx *app.Context) (interface{}, error) {
	username := ctx.Request.PathParam("username")

	return a.service.GetAccountWithBlogs(ctx, username)
}

// Update updates user details but not password.
func (a account) Update(ctx *app.Context) (interface{}, error) {
	user, err := ctx.Request.UnmarshalUser()
	if err != nil {
		return nil, err
	}

	return a.service.UpdateUser(ctx, user)
}

func (a account) UpdatePassword(ctx *app.Context) (interface{}, error) {
	pwd := struct {
		Old string `json:"old_password"`
		New string `json:"new_password"`
	}{}

	err := ctx.Request.Unmarshal(&pwd)
	if err != nil {
		return nil, err
	}

	return nil, a.service.UpdatePassword(ctx, pwd.Old, pwd.New)
}

func (a account) Delete(ctx *app.Context) (interface{}, error) {
	err := a.service.Delete(ctx)
	if err != nil {
		return nil, err
	}

	return a.Logout(ctx)
}

func (a account) CheckAvailability(ctx *app.Context) (interface{}, error) {
	username := ctx.Request.QueryParam("username")
	email := ctx.Request.QueryParam("email")
	phone := ctx.Request.QueryParam("phone")

	return nil, a.service.CheckAvailability(ctx, &models.User{UserName: username, PhoneNo: sql.NullString{String: phone, Valid: true}, Email: sql.NullString{String: email, Valid: true}})
}
