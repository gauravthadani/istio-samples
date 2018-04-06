package handler

import (
	"fmt"
	"log"
	"net/http"
)

func HelloWorld() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Service request")
		fmt.Fprintf(w, "Hello Universe !")
	})
}
