package main

import (
	_ "github.com/bmizerany/pq"
	_ "goweb/conversion"
	_ "goweb/model"
	_ "goweb/template"
	"goweb/view"
	"net/http"
)

func main() {
	http.HandleFunc("/", view.MakeHandler(view.HomeHandler))
	http.HandleFunc("/getjson", view.JSONHandler)
	//http.HandleFunc("/view/", MakeHandler(viewHandler))
	//http.HandleFunc("/edit/", MakeHandler(editHandler))
	//http.HandleFunc("/save/", MakeHandler(SaveHandler))
	http.ListenAndServe(":8080", nil)
}
