package main

import (
	"fmt"
	_ "gooo/introspection"
	_ "gooo/model"
	"gooo/router"
	"gooo/view"
	"net/http"
	"os"
)

var (
	addr     = ":8080"
	r        = router.New()
	assetDir = "static"
	pwd, _   = os.Getwd()
)

func main() {
	fmt.Printf("Gooo is now serving %s\n", addr)
	r.Get("/:foo/:bar", func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		foo := params.Get(":foo")
		bar := params.Get(":bar")
		fmt.Fprintf(w, "%s %s", foo, bar)
	})
	r.Get("/", view.PostHandler)
	r.Get("/posts", view.PostHandler)
	r.Post("/posts", view.NewPostHandler)
	r.Get("/hello/:first/:last", view.HelloHandler)
	r.Get("/setcookie", view.TestCookieSetHandler)
	r.Get("/getcookie", view.TestCookieGetHandler)
	r.Get("/count", view.CountHandler)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(assetDir))))
	http.Handle("/", r)
	http.ListenAndServe(addr, nil)
}
