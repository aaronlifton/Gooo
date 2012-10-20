package main

import (
	_ "gooo/introspection"
	_ "gooo/model"
	"gooo/view"
	"net/http"
)

func main() {
	http.HandleFunc("/", view.MakeHandler(view.HomeHandler))
	http.HandleFunc("/getjson", view.JSONHandler)
	http.ListenAndServe(":8080", nil)
}
