package controllers

import (
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
