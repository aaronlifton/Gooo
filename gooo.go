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

type Router struct {
}

func (p * Router) ServeHTTP (w http.ResponseWriter, r * http.Request) {
    switch r.URL.Path {
      case "/":
          view.HomeHandler(w,r)
          return
      case "/home":
           http.Redirect(w, r, "/", http.StatusFound)
           return
      case "/test":
          view.TestHandler(w,r)
          return
      case "/posts/new":
          view.NewPostHandler(w,r)
          return
      case "/hello":
          view.HelloHandler(w, r)
          return
    }
    http.NotFound(w, r)
    return
}

func main () {
	  fmt.Printf("Gooo is now serving %s\n", addr)
    mux := &Router{}
    http.ListenAndServe (addr, mux)
}

