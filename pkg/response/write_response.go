package response

import (
	"encoding/json"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/errors"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/log"
	"net/http"
)

func WriteError(w http.ResponseWriter, err error, log log.Logger) {
	var errType string

	log.Error(err)

	switch err.(type) {
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
	case nil:
		return
	}

	e := struct {
		Msg  string `json:"error,omitempty"`
		Type string `json:"type,omitempty"`
	}{
		Msg:  err.Error(),
		Type: errType,
	}

	resp, _ := json.Marshal(e)

	w.Write(resp)
}

func SetHeader(w http.ResponseWriter, err error, log log.Logger) {
	switch err.(type) {
	case errors.DBError, errors.Error:
		w.WriteHeader(http.StatusInternalServerError)
	case errors.MissingParam, errors.InvalidParam:
		w.WriteHeader(http.StatusBadRequest)
	case errors.EntityNotFound:
		w.WriteHeader(http.StatusNotFound)
	case errors.AuthError:
		w.WriteHeader(http.StatusUnauthorized)
	case nil:
		w.WriteHeader(http.StatusOK)
		return
	}

	WriteError(w, err, log)
}
