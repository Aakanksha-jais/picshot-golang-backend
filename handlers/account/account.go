package account

import (
	"database/sql"
	"net/http"

	"github.com/Aakanksha-jais/picshot-golang-backend/handlers"

	"github.com/Aakanksha-jais/picshot-golang-backend/models"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/auth"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/app"

	"github.com/Aakanksha-jais/picshot-golang-backend/services"
)

type account struct {
	service services.Account
}

func (a account) SendOTP(ctx *app.Context) (interface{}, error) {
	phone := ctx.Request.PathParam("phone")

	return a.service.SendOTP(ctx, phone)
}

func (a account) VerifyPhone(ctx *app.Context) (interface{}, error) {
	phone := struct {
		URL string `json:"url"`
		SID string `json:"sid"`
		OTP string `json:"otp"`
	}{}

	err := ctx.Request.Unmarshal(&phone)
	if err != nil {
		return nil, err
	}

	return nil, a.service.VerifyPhone(ctx, phone.SID, phone.OTP, phone.URL)
}

func New(service services.Account) handlers.Account {
	return account{service: service}
}

func (a account) Login(ctx *app.Context) (interface{}, error) {
	user, err := ctx.Request.UnmarshalUser()
	if err != nil {
		return nil, err
	}

	account, err := a.service.Login(ctx, user)
	if err != nil {
		return nil, err
	}

	token, err := auth.CreateToken(account.ID)
	if err != nil {
		return nil, err
	}

	ctx.Debugf("token generated successfully")

	ctx.SetAuthHeader(token)

	return &account.User, ctx.SetStatus(http.StatusOK)
}

func (a account) Signup(ctx *app.Context) (interface{}, error) {
	user, err := ctx.Request.UnmarshalUser()
	if err != nil {
		return nil, err
	}

	ctx.Debugf("signup request for %v", user)

	account, err := a.service.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	token, err := auth.CreateToken(account.ID)
	if err != nil {
		return nil, err
	}

	ctx.Debugf("token generated successfully")

	ctx.SetAuthHeader(token)

	return &account.User, ctx.SetStatus(http.StatusCreated)
}

func (a account) Logout(ctx *app.Context) (interface{}, error) { //todo invalidate the auth token
	ctx.SetAuthHeader("")

	return nil, ctx.SetStatus(http.StatusOK)
}

func (a account) Get(ctx *app.Context) (interface{}, error) {
	jwtIDKey := auth.JWTContextKey("claims")
	id := ctx.Value(jwtIDKey).(*auth.Claims).UserID

	ctx.Debugf("user with id: %v is logged in", id)

	return a.service.GetByID(ctx, id)
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

func (a account) JWKSEndpoint(_ *app.Context) (interface{}, error) {
	return auth.JWKS{
		Keys: []auth.JWK{
			{
				ID:             "lVHu/0QjmU/ZFq8oxD9KYnDt6NA=",
				Algorithm:      "RS256",
				Type:           "RSA",
				Use:            "sig",
				Operations:     []string{"verify"},
				Modulus:        "hUtynIfJOCeZTFBO77Yh3JznOWKbgBpTglbWMhuPWZElz2BO_G9hqomTFnCkEfMQzTbGgeu2yZskfXcn5yO0vJ1RzCIcHe77mKh1iIGQ1AjgM1t7vUkiyQmf3IK0CH-UPbsMgPi98jSZ_cEs_iEzLrizG47vuO8xo1hbKNDI9WldqUIqHVz9YZZEOTm7aClLx1bkBqHyosnbq_oUDViWEavpwzsOkq02Fvcm5KblxYlgSyigDnxaA3V_nOtdpQSzSibbYYbIcxqYCIF-iJvOi0COj5irWvM0oNvI9YzAxmyNNt132dbeKn8SYmjukif2CiPICmgoOKB8hEfZ7DLU-Hrd4jMG_pJT87M8jQymK0nE6j-w1fA_HsmBLdb0Xfrrbxpe3JydP_XXK_1-50K9A8wmZaJTslgrlVOG5Iy6K_QXIHmQuwSZpcxPmmVx3L75quPFnuAAY89dUbT60EalyB88ZFvPMllyPdFVotYHJayVO0kZHh4ep4qsYBuOYL5X3HVzW8afhYgZ3E11AwCrbxy9JgSyhXaMMu8i6w33o7ZNVqMX-woLQO-X_IONX4xTn-9Qu0jOs9UthkfBXISp6BKq5kcRJ5wMd0jUfV8d6E7IsKGRKRBSN6-_WbsJoyguu9i4OEetTc-RsVGP1InfBERLODKQ5d7mcW6RkkD23Es",
				PublicExponent: "AQAB",
			},
		},
	}, nil
}
