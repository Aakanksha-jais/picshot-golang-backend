package account

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Aakanksha-jais/picshot-golang-backend/models"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/auth"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/configs"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/errors"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/log"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/response"
)

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
		response.New(errors.Error{Err: err, Type: "body-read-error", Message: "error in reading request body"}, nil).WriteHeader(w)

		return nil, err
	}

	return body, nil
}

func unmarshalUser(w http.ResponseWriter, body []byte, logger log.Logger) (*models.User, error) {
	var u struct {
		ID       int64   `json:"-"`                  // Unique User ID
		UserName string  `json:"user_name"`          // Username
		FName    string  `json:"f_name"`             // First Name
		LName    string  `json:"l_name"`             // Last Name
		Email    *string `json:"email,omitempty"`    // Email
		PhoneNo  *string `json:"phone_no,omitempty"` // Phone Number
		Password string  `json:"password,omitempty"` // Password
	}

	err := json.Unmarshal(body, &u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response.New(errors.Error{Err: err, Type: "unmarshal-error", Message: "invalid request body"}, nil).Write(w)

		return nil, err
	}

	user := models.User{
		ID:       u.ID,
		UserName: u.UserName,
		FName:    u.FName,
		LName:    u.LName,
		Password: u.Password,
	}

	if u.Email != nil {
		user.Email = sql.NullString{String: *u.Email, Valid: true}
	}

	if u.PhoneNo != nil {
		user.PhoneNo = sql.NullString{String: *u.PhoneNo, Valid: true}
	}

	return &user, nil
}
