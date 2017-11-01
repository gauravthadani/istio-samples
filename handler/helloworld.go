package handler

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func HelloWorld() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Service request")
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
}
