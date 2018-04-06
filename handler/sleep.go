package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func Sleep() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		seconds := r.URL.Query().Get("seconds")
		if seconds != "" {
			i, _ := strconv.Atoi(seconds)

			fmt.Printf("Sleeping for %d seconds", i)
			time.Sleep(time.Duration(i) * time.Second)
		}
		fmt.Fprintf(w, "OK")
	})
}
