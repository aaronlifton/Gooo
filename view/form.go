package view

import "net/http"

func HandleForm(r *http.Request) {
	r.ParseForm()
	for k, v := range r.Form {
		validate(k, v)
	}
}

func validate(k string, v interface{}) {
	//TODO: validation
}
