package controllers

import (
	"net/http"

	"github.com/epiq122/epiqlens/views"
)

type Users struct {
	Templates struct {
		Signup views.Template
	}
}

func (u Users) Signup(w http.ResponseWriter, r *http.Request) {
	u.Templates.Signup.Execute(w, nil)
}
