package controllers

import (
	"fmt"
	"net/http"
)

type Users struct {
	Templates struct {
		Signup Template
	}
}

func (u Users) Signup(w http.ResponseWriter, r *http.Request) {
	u.Templates.Signup.Execute(w, nil)
}

func (u Users) CreateUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "<p>Email: %s</p>", r.FormValue("email"))
	fmt.Fprintf(w, "<p>Password: %s</p>", r.FormValue("password"))
}
