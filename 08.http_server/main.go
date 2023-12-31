package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Home, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Login, %s, %s", r.URL.Query().Get("username"), r.URL.Query().Get("password"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
