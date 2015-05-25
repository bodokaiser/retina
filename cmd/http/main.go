package main

import "flag"
import "net/http"

import "github.com/bodokaiser/retina/httpd"

func main() {
	listen := *flag.String("listen", ":3000", "")

	http.HandleFunc("/", httpd.NewHandler().ServeHTTP)
	http.ListenAndServe(listen, nil)
}
