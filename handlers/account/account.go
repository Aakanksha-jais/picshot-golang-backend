package account

import (
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

func (a account) LogIn(w http.ResponseWriter, r *http.Request) {
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

	account, err := a.service.Get(r.Context(), user)
	if err != nil {
		response.SetHeader(w, err, a.logger)

		return
	}

	token, err := generateToken(expirationTime, account.ID, a.config)
	if err != nil {
		a.logger.Errorf("error in generating token: %v", err)
	} else {
		a.logger.Infof("token generated successfully: %v, expires at: %v", token[:10], expirationTime.Format(time.RFC850))
		setAuthHeader(w, token)
	}

	response.SetHeader(w, err, a.logger) // Set Header to StatusOK if err is nil
}

func (a account) SignUp(w http.ResponseWriter, r *http.Request) {
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

	// Create an Account based on User Details Provided
	account, err := a.service.Create(r.Context(), user)
	if err != nil {
		response.SetHeader(w, err, a.logger)

		return
	}

	// Create a JWT Token
	token, err := generateToken(expirationTime, account.ID, a.config)
	if err != nil {
		a.logger.Errorf("error in generating token: %v", err)
		response.SetHeader(w, err, a.logger)

		return
	}

	a.logger.Infof("token generated successfully: %v, expiration: %v", token[:10], expirationTime)

	setAuthHeader(w, token)

	w.WriteHeader(http.StatusCreated)
}

func setAuthHeader(w http.ResponseWriter, token string) {
	w.Header().Add("Authorization", fmt.Sprintf("bearer %s", token))
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
		response.SetHeader(w, errors.Error{Err: err, Type: "body-read-error", Message: "error in reading request body"}, logger)

		return nil, err
	}

	return body, nil
}

func unmarshalUser(w http.ResponseWriter, body []byte, logger log.Logger) (*models.User, error) {
	var user models.User

	err := json.Unmarshal(body, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response.WriteError(w, errors.Error{Err: err, Type: "unmarshal-error", Message: "invalid request body"}, logger)

		return nil, err
	}

	return &user, nil
}

func (a account) LogOut(w http.ResponseWriter, r *http.Request) {
}
