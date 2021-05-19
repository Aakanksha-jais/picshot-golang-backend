package handlers

import (
	"net/http"
)

type Account interface {
	Login(w http.ResponseWriter, r *http.Request)
	Signup(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	CheckAvailability(w http.ResponseWriter, r *http.Request)
}

type Blog interface {
	GetAll(w http.ResponseWriter, r *http.Request)
}
