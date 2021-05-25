package app

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"mime/multipart"
	"net/http"

	"github.com/Aakanksha-jais/picshot-golang-backend/models"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/errors"

	"github.com/gorilla/mux"
)

type Request struct {
	req        *http.Request
	pathParams map[string]string
}

func NewRequest(r *http.Request) *Request {
	return &Request{
		req:        r,
		pathParams: mux.Vars(r),
	}
}

func (r *Request) Context() context.Context {
	return r.req.Context()
}

func (r *Request) URL() string {
	return r.req.URL.Path
}

func (r *Request) QueryParam(key string) string {
	return r.req.URL.Query().Get(key)
}

func (r *Request) Header(key string) string {
	return r.req.Header.Get(key)
}

func (r *Request) PathParam(key string) string {
	return r.pathParams[key]
}

const maxSize = int64(32 << 20) // max 32 MB size

func (r *Request) ParseImages() []*multipart.FileHeader {
	err := r.req.ParseMultipartForm(maxSize)
	if err != nil {
		return nil
	}

	return r.req.MultipartForm.File["image"]
}

func (r *Request) FormValue(key string) string {
	return r.req.FormValue(key)
}

func (r *Request) Unmarshal(i interface{}) error {
	body, err := r.body()
	if err != nil {
		return errors.BodyRead{Err: err}
	}

	err = json.Unmarshal(body, i)
	if err != nil {
		return errors.Unmarshal{Err: err}
	}

	return nil
}

func (r *Request) body() ([]byte, error) {
	bodyBytes, err := ioutil.ReadAll(r.req.Body)
	if err != nil {
		return nil, errors.BodyRead{Err: err}
	}

	r.req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	return bodyBytes, nil
}

func (r *Request) UnmarshalUser() (*models.User, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	var u struct {
		UserName string  `json:"user_name"`          // Username
		FName    string  `json:"f_name"`             // First Name
		LName    string  `json:"l_name"`             // Last Name
		Email    *string `json:"email,omitempty"`    // Email
		PhoneNo  *string `json:"phone_no,omitempty"` // Phone Number
		Password string  `json:"password,omitempty"` // Password
	}

	err = json.Unmarshal(body, &u)
	if err != nil {
		return nil, errors.Unmarshal{Err: err}
	}

	user := &models.User{
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

	return user, nil
}
