package main

import "flag"
import "net/http"

import "github.com/bodokaiser/retina/httpd"

type config struct {
	address string
}

func main() {
	c := flags()

	http.HandleFunc("/", httpd.NewHandler().ServeHTTP)
	http.ListenAndServe(c.address, nil)
}

func flags() config {
	c := config{}

	flag.StringVar(&c.address, "listen", ":3000", "")
	flag.Parse()

	return c
}
