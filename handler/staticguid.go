package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/satori/go.uuid"
)

func StaticGuid() http.Handler {

	u2, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return nil
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("static guid handler")
		fmt.Fprintf(w, "%s", u2.String())
	})
}
