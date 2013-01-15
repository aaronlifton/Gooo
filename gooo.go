package main

import (
	"fmt"
	_ "gooo/introspection"
	_ "gooo/model"
	"gooo/view"
	"net/http"
)

var (
	addr = ":8080"
)

func main() {
	fmt.Printf("Gooo is now serving %s\n", addr)
	http.HandleFunc("/", view.MakeHandler(view.HomeHandler))
	http.HandleFunc("/test", view.MakeHandler(view.TestHandler))
	http.HandleFunc("/getjson", view.JSONHandler)
	http.ListenAndServe(addr, nil)
}
