package main

import (
	_ "github.com/bmizerany/pq"
	_ "goweb/conversion"
	_ "goweb/model"
	_ "goweb/template"
	"goweb/view"
	"net/http"
)

const dbParams string = `host=ec2-54-243-239-221.compute-1.amazonaws.com user=kwcwqwdgfelhrs port=5432 password=PASSWORD_HERE dbname=d62335du1mgsdc sslmode=require`

func main() {
	http.HandleFunc("/", view.MakeHandler(view.HomeHandler))
	http.HandleFunc("/getjson", view.JSONHandler)
	//http.HandleFunc("/view/", MakeHandler(viewHandler))
	//http.HandleFunc("/edit/", MakeHandler(editHandler))
	//http.HandleFunc("/save/", MakeHandler(SaveHandler))
	http.ListenAndServe(":8080", nil)
}
