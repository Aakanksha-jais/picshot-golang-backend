package account

import (
	"encoding/json"
	"github.com/Aakanksha-jais/picshot-golang-backend/handlers"
	"github.com/Aakanksha-jais/picshot-golang-backend/models"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/auth"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/configs"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/errors"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/log"
	"github.com/Aakanksha-jais/picshot-golang-backend/services"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type account struct {
	service services.Account
	logger  log.Logger
	config  configs.ConfigLoader
}

func New(service services.Account, logger log.Logger, config configs.ConfigLoader) handlers.Account {
	return account{
		service: service,
		logger:  logger,
		config:  config,
	}
}

func (a account) LogIn(w http.ResponseWriter, r *http.Request) {
	expirationTime := time.Now().Add(5 * time.Minute)

	body, err := readUser(w, r.Body)
	if err != nil {
		a.logger.Errorf("error in reading request body: %v", err)
		return
	}

	user, err := unmarshalUser(w, body)
	if err != nil {
		a.logger.Errorf("error in unmarshalling request body %v", err)
		return
	}

	account, err := a.service.Get(r.Context(), user)
	if err != nil {
		handlers.SetHeader(w, err)
	}

	// Create a JWT Token
	token, err := generateToken(expirationTime, account.ID, a.config)
	handlers.SetHeader(w, err) // Set Header to StatusOK if err is nil

	if err != nil {
		return
	}

	setCookie(w, token, expirationTime)
}

func (a account) SignUp(w http.ResponseWriter, r *http.Request) {
	expirationTime := time.Now().Add(5 * time.Minute)

	body, err := readUser(w, r.Body)
	if err != nil {
		a.logger.Errorf("error in reading request body: %v", err)
		return
	}

	user, err := unmarshalUser(w, body)
	if err != nil {
		a.logger.Errorf("error in unmarshalling request body %v", err)
		return
	}

	// Create an Account based on User Details Provided
	account, err := a.service.Create(r.Context(), user)
	if err != nil {
		handlers.SetHeader(w, err)
	}

	// Create a JWT Token
	token, err := generateToken(expirationTime, account.ID, a.config)
	if err != nil {
		handlers.SetHeader(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)

	setCookie(w, token, expirationTime)
}

func setCookie(w http.ResponseWriter, token string, exp time.Time) {
	http.SetCookie(w, &http.Cookie{
		Name:     "auth-token",
		Value:    token,
		Expires:  exp,
		HttpOnly: true,
		Secure:   true,
	})
}

func generateToken(exp time.Time, userID int64, config configs.ConfigLoader) (string, error) {
	token, err := auth.CreateToken(config, auth.NewClaim(exp.Unix(), userID))
	if err != nil {
		return "", err
	}

	return token, nil
}

func readUser(w http.ResponseWriter, reqBody io.ReadCloser) ([]byte, error) {
	body, err := ioutil.ReadAll(reqBody)
	if err != nil {
		handlers.SetHeader(w, errors.Error{Err: err, Type: "body-read-error", Message: "error in reading request body"})

		return nil, err
	}

	return body, nil
}

func unmarshalUser(w http.ResponseWriter, body []byte) (*models.User, error) {
	var user models.User

	err := json.Unmarshal(body, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		handlers.WriteError(w, errors.Error{Err: err, Type: "unmarshal-error", Message: "invalid request body"})

		return nil, err
	}

	return &user, nil
}

func (a account) LogOut(w http.ResponseWriter, r *http.Request) {
}
