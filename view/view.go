package view

import (
	"bytes"
	"gooo/introspection"
	"gooo/model"
	"html/template"
	"net/http"
	"time"
)

// template config
var (
	templateCache          = false
	templatePattern string = "tmpl/*.html"
	templates       *template.Template
)

func ParseTemplateGlob(pattern string, cache bool) {
	templateCache = cache
	templatePattern = pattern
	templates = template.Must(template.ParseGlob(templatePattern))
}

func RenderTemplate(w http.ResponseWriter, tmpl string, context map[string]interface{}) {
	var err error
	var buf bytes.Buffer

	// you can disable template caching to speedup
	// the development process. in this case we
	// always re-parse all templates
	if templateCache == false {
		tmp_templates := template.Must(template.ParseGlob(templatePattern))
		err = tmp_templates.ExecuteTemplate(&buf, tmpl, context)
		// in production mode, use the cached template
		// and load the template by name
	} else {
		err = templates.ExecuteTemplate(&buf, tmpl, context)
	}
	// on error serve the internal service error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// set the content length, type, etc.
	w.Write(buf.Bytes())
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	db := model.OpenConn()
	//model.TestEmptyDB()

	var p model.Post = model.Post{2, "Hello World", "whats up yo", 1, true, time.Now(), time.Now()}
	var p2 model.Post = model.Post{2, "Test2", "another test post please ignore", 1, true, time.Now(), time.Now()}
	atts := introspection.GetStructValues(&p)
	model.InsertIntoDB(atts)
	//model.GetPosts(10)
	posts := map[string]interface{}{"p1": introspection.ConvertToMap(p), "p2": introspection.ConvertToMap(p2)}
	ctx := map[string]interface{}{"posts": posts}
	defer db.Close()
	RenderTemplate(w, "index", ctx)
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	var p model.Post = model.Post{2, "Hello World", "whats up yo", 1, true, time.Now(), time.Now()}
	var p2 model.Post = model.Post{2, "Test2", "another test post please ignore", 1, true, time.Now(), time.Now()}
	//atts := introspection.GetStructValues(&p)
	posts := map[string]interface{}{"p1": introspection.ConvertToMap(p), "p2": introspection.ConvertToMap(p2)}
	ctx := map[string]interface{}{"posts": posts}
	RenderTemplate(w, "index", ctx)
}

func SaveHandler(w http.ResponseWriter, r *http.Request) {
	content := r.FormValue("content")
	p := &model.Post{Title: "lol", Content: content, UserId: 1, Published: true, Created: time.Now(), Modified: time.Now()}
	atts := introspection.GetStructValues(&p)
	model.InsertIntoDB(atts)
	ctx := map[string]interface{}{"post": introspection.ConvertToMap(p)}
	RenderTemplate(w, "post", ctx)
}

func MakeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)
	}
}

func JSONHandler(w http.ResponseWriter, r *http.Request) {
	var p model.Post = model.Post{2, "Test", "test post please ignore", 1, true, time.Now(), time.Now()}
	w.Header().Set("Content-Type", "application/json")
	b := introspection.ConvertToJson(p)
	w.Write(b)
	//fmt.Fprintf(w, renderJson(w, res))
}
