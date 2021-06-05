package app

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Aakanksha-jais/picshot-golang-backend/models"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/errors"
)

type Handler func(c *Context) (interface{}, error)

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx, _ := r.Context().Value(appContextKey).(*Context)
	data, err := h(ctx)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var (
		respData interface{}
		resp     interface{}
		errtype  string
	)

	if _, ok := data.(func(w http.ResponseWriter)); !ok {
		errtype = setHeader(w, err)
	}

	switch data := data.(type) {
	case *models.Account:
		respData = getAccountResponse(data)
	case func(w http.ResponseWriter):
		data(w)
	case *models.User:
		respData = getUserResponse(data)
	default:
		respData = data
	}

	switch err {
	case nil:
		resp = struct {
			Status string      `json:"status"`
			Data   interface{} `json:"data,omitempty"`
		}{
			Status: "success",
			Data:   respData,
		}

	default:
		type respError struct {
			Msg  string `json:"msg,omitempty"`
			Type string `json:"type,omitempty"`
		}

		resp = struct {
			Status string    `json:"status"`
			Err    respError `json:"error"`
		}{
			Status: "failure",
			Err: respError{
				Msg:  err.Error(),
				Type: errtype,
			},
		}

		ctx.Error(err)
	}

	response, err := json.Marshal(resp)
	if err != nil {
		ctx.Error(err)
	}

	_, err = w.Write(response)
	if err != nil {
		ctx.Error(err)
	}
}

func setHeader(w http.ResponseWriter, err error) string {
	switch err := err.(type) {
	case nil:
		w.WriteHeader(http.StatusOK)
	case errors.DBError:
		w.WriteHeader(http.StatusInternalServerError)
		return "db-error"
	case errors.MissingParam:
		w.WriteHeader(http.StatusBadRequest)
		return "missing-param"
	case errors.InvalidParam:
		w.WriteHeader(http.StatusBadRequest)
		return "invalid-param"
	case errors.EntityNotFound:
		w.WriteHeader(http.StatusNotFound)
		return "entity-not-found"
	case errors.AuthError:
		w.WriteHeader(http.StatusUnauthorized)
		return "auth-error"
	case errors.Error:
		w.WriteHeader(http.StatusInternalServerError)
		return err.Type
	case errors.EntityAlreadyExists:
		w.WriteHeader(http.StatusBadRequest)
		return "entity-already-exists"
	case errors.Unmarshal:
		w.WriteHeader(http.StatusBadRequest)
		return "unmarshal-error"
	case errors.BodyRead:
		w.WriteHeader(http.StatusBadRequest)
		return "body-read-error"
	}

	return ""
}

type userResp struct {
	UserName string  `json:"user_name"`
	FName    string  `json:"f_name"`
	LName    string  `json:"l_name"`
	Email    *string `json:"email,omitempty"`
	PhoneNo  *string `json:"phone_no,omitempty"`
}

func getUserResponse(data interface{}) userResp {
	user, _ := data.(*models.User)
	if user == nil {
		return userResp{}
	}

	u := struct {
		UserName string  `json:"user_name"`
		FName    string  `json:"f_name"`
		LName    string  `json:"l_name"`
		Email    *string `json:"email,omitempty"`
		PhoneNo  *string `json:"phone_no,omitempty"`
	}{
		UserName: user.UserName,
		FName:    user.FName,
		LName:    user.LName,
	}

	if user.Email.String != "" {
		u.Email = &user.Email.String
	}

	if user.PhoneNo.String != "" {
		u.PhoneNo = &user.PhoneNo.String
	}

	return u
}

type accountResp struct {
	UserName   string        `json:"user_name"`
	FName      string        `json:"f_name"`
	LName      string        `json:"l_name"`
	Email      *string       `json:"email,omitempty"`
	PhoneNo    *string       `json:"phone_no,omitempty"`
	PwdUpdate  *time.Time    `json:"pwd_update,omitempty"`
	Blogs      []models.Blog `json:"blogs,omitempty"`
	CreatedAt  time.Time     `json:"created_at"`
	DelRequest *time.Time    `json:"del_req,omitempty"`
	Status     string        `json:"status"`
}

func getAccountResponse(data interface{}) accountResp {
	account, _ := data.(*models.Account)
	if account == nil {
		return accountResp{}
	}

	a := accountResp{
		UserName:   account.UserName,
		FName:      account.FName,
		LName:      account.LName,
		Blogs:      account.Blogs,
		CreatedAt:  account.CreatedAt,
		DelRequest: nil,
		Status:     account.Status,
	}

	if account.Email.String != "" {
		a.Email = &account.Email.String
	}

	if account.PhoneNo.String != "" {
		a.PhoneNo = &account.PhoneNo.String
	}

	if account.PwdUpdate.Valid {
		a.PwdUpdate = &account.PwdUpdate.Time
	}

	if account.DelRequest.Valid {
		a.DelRequest = &account.DelRequest.Time
	}

	return a
}
