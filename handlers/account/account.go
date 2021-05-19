package account

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/Aakanksha-jais/picshot-golang-backend/handlers"
	"github.com/Aakanksha-jais/picshot-golang-backend/models"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/auth"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/configs"
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
		response.New(err, nil).WriteHeader(w)

		return
	}

	token, err := generateToken(expirationTime, account.ID, a.config)
	if err != nil {
		a.logger.Errorf("error in generating token: %v", err)
	} else {
		a.logger.Infof("token generated successfully: %v, expires at: %v", token[:10], expirationTime.Format(time.RFC850))
		setAuthHeader(w, token)
	}

	response.New(err, nil).WriteHeader(w) // Set Header to StatusOK if err is nil
}

func (a account) CheckAvailability(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	email := r.URL.Query().Get("email")
	phone := r.URL.Query().Get("phone")

	err := a.service.CheckAvailability(r.Context(), &models.User{UserName: username, PhoneNo: sql.NullString{String: phone, Valid: true}, Email: sql.NullString{String: email, Valid: true}})
	response.New(err, nil).WriteHeader(w)
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
		response.New(err, nil).WriteHeader(w)

		return
	}

	// Create a JWT Token
	token, err := generateToken(expirationTime, account.ID, a.config)
	if err != nil {
		a.logger.Errorf("error in generating token: %v", err)
		response.New(err, nil).WriteHeader(w)

		return
	}

	a.logger.Infof("token generated successfully: %v, expiration: %v", token[:10], expirationTime.Format(time.RFC850))

	setAuthHeader(w, token)
	w.WriteHeader(http.StatusCreated)

	response.New(nil, nil).Write(w)
}

func (a account) Get(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(auth.JWTContextKey("user_id"))
	a.logger.Debugf("user with id: %v is logged in", id)

	account, err := a.service.GetByID(r.Context(), id.(int64))

	response.New(err, account).WriteHeader(w)
}

func (a account) GetUser(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]

	account, err := a.service.GetAccountWithBlogs(r.Context(), username)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	response.New(err, account).Write(w)
}

func (a account) Logout(w http.ResponseWriter, r *http.Request) {
}

func (a account) Update(w http.ResponseWriter, r *http.Request) {
}
