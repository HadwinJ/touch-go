package controllers

import "net/http"

// RegisterController function
func RegisterController() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
