package controllers

import (
	"html/template"
	"net/http"

	"github.com/epiq122/epiqlens/views"
)

type Static struct {
	Template Template
}

func (static Static) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	static.Template.Execute(w, nil)
}

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl views.Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "Is there going to be a free version?",
			Answer:   "Yes! We will be offering a 30 day free trial",
		},
		{
			Question: "What are your support hours?",
			Answer:   "These are TBD",
		},
		{
			Question: "How do I contact support?",
			Answer:   `Email us - <a href="mailto:epiqlens@gmail.com">epiqlens@gmail.com</a>`,
		},
		{
			Question: "When is the release?",
			Answer:   "Christmas 2024",
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
