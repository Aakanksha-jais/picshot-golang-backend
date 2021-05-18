package account

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Aakanksha-jais/picshot-golang-backend/handlers"
	"github.com/Aakanksha-jais/picshot-golang-backend/models"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/auth"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/configs"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/errors"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/log"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/response"
	"github.com/Aakanksha-jais/picshot-golang-backend/services"
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

func (a account) Login(w http.ResponseWriter, r *http.Request) {
	expirationTime := time.Now().Add(30 * time.Minute)

	body, err := readUser(w, r.Body, a.logger)
	if err != nil {
		a.logger.Errorf("error in reading request body: %v", err)
		return
	}

	user, err := unmarshalUser(w, body, a.logger)
	if err != nil {
		a.logger.Errorf("error in unmarshalling request body %v", err)
		return
	}

	account, err := a.service.Login(r.Context(), user)
	if err != nil {
		response.SetHeader(w, err, nil, a.logger)

		return
	}

	token, err := generateToken(expirationTime, account.ID, a.config)
	if err != nil {
		a.logger.Errorf("error in generating token: %v", err)
	} else {
		a.logger.Infof("token generated successfully: %v, expires at: %v", token[:10], expirationTime.Format(time.RFC850))
		setAuthHeader(w, token)
	}

	response.SetHeader(w, err, nil, a.logger) // Set Header to StatusOK if err is nil
}

func (a account) CheckAvailability(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	email := r.URL.Query().Get("email")
	phone := r.URL.Query().Get("phone")

	err := a.service.CheckAvailability(r.Context(), models.User{UserName: username, PhoneNo: sql.NullString{String: phone, Valid: true}, Email: sql.NullString{String: email, Valid: true}})
	response.SetHeader(w, err, nil, a.logger)
}

func (a account) Signup(w http.ResponseWriter, r *http.Request) {
	expirationTime := time.Now().Add(30 * time.Minute)

	body, err := readUser(w, r.Body, a.logger)
	if err != nil {
		a.logger.Errorf("error in reading request body: %v", err)
		return
	}

	user, err := unmarshalUser(w, body, a.logger)
	if err != nil {
		a.logger.Errorf("error in unmarshalling request body %v", err)
		return
	}

	a.logger.Debugf("signup request for %v", user)

	// Create an Account based on User Details Provided
	account, err := a.service.Create(r.Context(), user)
	if err != nil {
		response.SetHeader(w, err, nil, a.logger)

		return
	}

	// Create a JWT Token
	token, err := generateToken(expirationTime, account.ID, a.config)
	if err != nil {
		a.logger.Errorf("error in generating token: %v", err)
		response.SetHeader(w, err, nil, a.logger)

		return
	}

	a.logger.Infof("token generated successfully: %v, expiration: %v", token[:10], expirationTime.Format(time.RFC850))

	setAuthHeader(w, token)

	w.WriteHeader(http.StatusCreated)
	response.WriteResponse(w, nil, nil, a.logger)
}

func (a account) Get(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(auth.JWTContextKey("user_id"))
	a.logger.Debugf("user with id: %v is logged in", id)

	account, err := a.service.GetByID(r.Context(), id.(int64))

	response.SetHeader(w, err, account, a.logger)
}

func setAuthHeader(w http.ResponseWriter, token string) {
	w.Header().Add("Authorization", fmt.Sprintf("Bearer %s", token))
}

func generateToken(exp time.Time, userID int64, config configs.ConfigLoader) (string, error) {
	token, err := auth.CreateToken(config, auth.NewClaim(exp.Unix(), userID))
	if err != nil {
		return "", err
	}

	return token, nil
}

func readUser(w http.ResponseWriter, reqBody io.ReadCloser, logger log.Logger) ([]byte, error) {
	body, err := ioutil.ReadAll(reqBody)
	if err != nil {
		response.SetHeader(w, errors.Error{Err: err, Type: "body-read-error", Message: "error in reading request body"}, nil, logger)

		return nil, err
	}

	return body, nil
}

func unmarshalUser(w http.ResponseWriter, body []byte, logger log.Logger) (*models.User, error) {
	var user models.User

	err := json.Unmarshal(body, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response.WriteResponse(w, errors.Error{Err: err, Type: "unmarshal-error", Message: "invalid request body"}, nil, logger)

		return nil, err
	}

	return &user, nil
}

func (a account) Logout(w http.ResponseWriter, r *http.Request) {
}

func (a account) Update(w http.ResponseWriter, r *http.Request) {
}
