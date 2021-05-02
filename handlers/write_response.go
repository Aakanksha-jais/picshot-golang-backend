package handlers

import (
	"encoding/json"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/errors"
	"net/http"
)

func WriteError(w http.ResponseWriter, err error) {
	var errType string

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

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

func SetHeader(w http.ResponseWriter, err error) {
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

	WriteError(w, err)
}
