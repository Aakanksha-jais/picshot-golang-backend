package response

import (
	"encoding/json"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/errors"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/log"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, err error, resp interface{}, log log.Logger) {
	var errType string

	switch err.(type) {
	case nil:
	case errors.DBError:
		errType = "db-error"
	case errors.MissingParam:
		errType = "missing-param"
	case errors.InvalidParam:
		errType = "invalid-param"
	case errors.EntityNotFound:
		errType = "entity-not-found"
	case errors.AuthError:
		errType = "auth-error"
	case errors.Error:
		errType = err.(errors.Error).Type
	case errors.EntityAlreadyExists:
		errType = "entity-already-exists"
	}

	var result interface{}

	if err == nil {
		result = struct {
			Status string      `json:"status"`
			Data   interface{} `json:"data,omitempty"`
		}{
			Status: "success",
			Data:   resp,
		}
	} else {
		type Err struct {
			Msg  string `json:"msg,omitempty"`
			Type string `json:"type,omitempty"`
		}
		result = struct {
			Status string `json:"status"`
			Err    Err    `json:"error"`
		}{
			Status: "failure",
			Err: Err{
				Msg:  err.Error(),
				Type: errType,
			},
		}

		log.Error(err)
	}

	response, _ := json.Marshal(result)

	w.Write(response)
}

func SetHeader(w http.ResponseWriter, err error, resp interface{}, log log.Logger) {
	switch err.(type) {
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

	WriteResponse(w, err, resp, log)
}
