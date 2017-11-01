package main

import (
	"log"
	"net/http"

	"github.com/gauravthadani/istio-samples/handler"
)

func main() {
	log.Println("Starting service")

	http.Handle("/hello", handler.HelloWorld())

	log.Fatal(http.ListenAndServe(":8080", nil))
}
