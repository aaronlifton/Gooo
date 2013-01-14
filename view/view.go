package view

import (
	"bytes"
	"fmt"
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

func HomeHandler(w http.ResponseWriter, r *http.Request, title string) {
	db := model.OpenConn()
	model.TestEmptyDB(db)
	var p model.Post = model.Post{model.M, 2, "Hello World", "whats up yo", 1, true, time.Now(), time.Now()}
	var p2 model.Post = model.Post{model.M, 2, "Test2", "another test post please ignore", 1, true, time.Now(), time.Now()}
	fmt.Println(p.ModelName())
	atts := introspection.GetStructValues(&p)
	model.InsertIntoDB(atts)
	model.GetPosts(10)
	posts := map[string]interface{}{"p1": introspection.ConvertToMap(p), "p2": introspection.ConvertToMap(p2)}
	ctx := map[string]interface{}{"posts": posts}
	defer db.Close()
	//defer stmt.Close()
	RenderTemplate(w, "index", ctx)
}

func SaveHandler(w http.ResponseWriter, r *http.Request, title string) {
	content := r.FormValue("content")
	p := &model.Post{Title: title, Content: content, UserId: 1, Published: true, Created: time.Now(), Modified: time.Now()}
	p.Title = "lol"
	atts := introspection.GetStructValues(&p)
	model.InsertIntoDB(atts)
	ctx := map[string]interface{}{"post": introspection.ConvertToMap(p)}
	RenderTemplate(w, "post", ctx)
}

func MakeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title := "title here"
		fn(w, r, title)
	}
}

func JSONHandler(w http.ResponseWriter, r *http.Request) {
	var p model.Post = model.Post{model.M, 2, "Test", "test post please ignore", 1, true, time.Now(), time.Now()}
	w.Header().Set("Content-Type", "application/json")
	b := introspection.ConvertToJson(p)
	w.Write(b)
	//fmt.Fprintf(w, renderJson(w, res))
}
