package response

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Aakanksha-jais/picshot-golang-backend/models"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/errors"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/log"
)

type Response struct {
	Err  error
	Data interface{}
}

func New(err error, data interface{}) Response {
	return Response{Err: err, Data: data}
}

func (r Response) Write(w http.ResponseWriter) {
	errType := checkErrorType(r.Err)

	var result interface{}

	if r.Err == nil {
		if account, ok := r.Data.(*models.Account); ok {
			a := struct {
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
			}{
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

			if account.PwdUpdate != nil {
				a.PwdUpdate = &account.PwdUpdate.Time
			}

			if account.DelRequest != nil {
				a.DelRequest = &account.DelRequest.Time
			}

			r.Data = a
		}

		if user, ok := r.Data.(*models.User); ok {
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

			r.Data = u
		}

		result = struct {
			Status string      `json:"status"`
			Data   interface{} `json:"data,omitempty"`
		}{
			Status: "success",
			Data:   r.Data,
		}
	} else {
		type err struct {
			Msg  string `json:"msg,omitempty"`
			Type string `json:"type,omitempty"`
		}
		result = struct {
			Status string `json:"status"`
			Err    err    `json:"error"`
		}{
			Status: "failure",
			Err: err{
				Msg:  r.Err.Error(),
				Type: errType,
			},
		}

		log.NewLogger().Error(r.Err) //todo
	}

	response, _ := json.Marshal(result)

	w.Write(response)
}

func checkErrorType(err error) string {
	switch err.(type) {
	case nil:
	case errors.DBError:
		return "db-error"
	case errors.MissingParam:
		return "missing-param"
	case errors.InvalidParam:
		return "invalid-param"
	case errors.EntityNotFound:
		return "entity-not-found"
	case errors.AuthError:
		return "auth-error"
	case errors.Error:
		return err.(errors.Error).Type
	case errors.EntityAlreadyExists:
		return "entity-already-exists"
	}

	return "unknown-type"
}

func (r Response) WriteHeader(w http.ResponseWriter) {
	switch r.Err.(type) {
	case errors.DBError, errors.Error:
		w.WriteHeader(http.StatusInternalServerError)
	case errors.MissingParam, errors.InvalidParam, errors.EntityAlreadyExists:
		w.WriteHeader(http.StatusBadRequest)
	case errors.EntityNotFound:
		w.WriteHeader(http.StatusNotFound)
	case errors.AuthError:
		w.WriteHeader(http.StatusUnauthorized)
	case nil:
		w.WriteHeader(http.StatusOK)
	}

	r.Write(w)
}
