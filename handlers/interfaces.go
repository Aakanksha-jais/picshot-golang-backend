package handlers

import "net/http"

type Account interface {
	LogIn(w http.ResponseWriter, r *http.Request)
	SignUp(w http.ResponseWriter, r *http.Request)
	LogOut(w http.ResponseWriter, r *http.Request)
	CheckAvailability(w http.ResponseWriter, r *http.Request)
}

type Blog interface {
	GetAll(w http.ResponseWriter, r *http.Request)
}
