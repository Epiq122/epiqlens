package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, "<h1>Epiq lens</h1")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch with me email me at <a href=\"mailto:epiqlens@gmail.com\">epiqlens@gmail.com</a>.</p>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.HandleFunc("/contact", contactHandler)
	fmt.Println("Starting server on :3000")
	http.ListenAndServe(":3000", nil)

}
