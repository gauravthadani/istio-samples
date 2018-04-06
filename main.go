package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gauravthadani/istio-samples/handler"
)

func report(r *http.Request) {
	r.Host = "httpbin.org"
	r.URL.Host = r.Host
	r.URL.Scheme = "http"
}

func main() {
	log.Println("Starting service")
	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   "httpbin.org",
	})
	proxy.Director = report
	http.Handle("/", proxy)
	http.Handle("/hello", handler.HelloWorld())
	http.Handle("/alternate", handler.Alternate())
	http.Handle("/unique", handler.StaticGuid())
	http.Handle("/sleep", handler.Sleep())

	log.Fatal(http.ListenAndServe(":8080", nil))
}
