package handler

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

var val bool = true

func Alternate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("alternate")

		if val == true {
			val = !val
			fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
			return
		}
		val = !val
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
	})
}
