package view

import (
	"bytes"
  "fmt"
	"gooo/introspection"
	"gooo/model"
	"html/template"
	"net/http"
	"time"
  "regexp"
  "strconv"
  "sync"
)

// template config
var (
	templateCache          = false
	templatePattern string = "tmpl/*.html"
	templates       *template.Template
)

type m map[string]interface{}



type Route struct {
  explicit bool
  h Handler
  re *regexp.Regexp
}

type Router struct {
  mu sync.RWMutex // lock for req involving concurrent processing
  m map[string]Route // Routing rules
}

type Handler interface {
  ServeHTTP(http.ResponseWriter, *http.Request)
}

//mine
func ParseTemplateGlob(pattern string, cache bool) {
	templateCache = cache
	templatePattern = pattern
	templates = template.Must(template.ParseGlob(templatePattern))
}

func RenderTemplate(w http.ResponseWriter, tmpl string, context m) {
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
  latestPosts := model.GetPosts(10)
	ctx := m{"posts": latestPosts}
	defer db.Close()
	RenderTemplate(w, "index", ctx)
}

func NewPostHandler(w http.ResponseWriter, r *http.Request) {
  title := r.FormValue("title")
  body := r.FormValue("content")
  userId, err := strconv.Atoi(r.FormValue("userId"))
  if err != nil {
    userId = 0
  }
  published := true
  p := model.Post{0,title, body, userId, published, time.Now(), time.Now()}
	atts := introspection.GetStructValues(&p)
	model.InsertIntoDB(atts)
  http.Redirect(w, r, "/", http.StatusFound)
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	var p model.Post = model.Post{0, "Hello World", "whats up yo", 1, true, time.Now(), time.Now()}
	var p2 model.Post = model.Post{0, "Test2", "another test post please ignore", 1, true, time.Now(), time.Now()}
	//atts := introspection.GetStructValues(&p)
	posts := m{"p1": introspection.ConvertToMap(p), "p2": introspection.ConvertToMap(p2)}
	ctx := m{"posts": posts}
	RenderTemplate(w, "index", ctx)
}

func MakeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)
	}
}

func JSONHandler(w http.ResponseWriter, r *http.Request) {
	var p model.Post = model.Post{0, "Test", "test post please ignore", 1, true, time.Now(), time.Now()}
	w.Header().Set("Content-Type", "application/json")
	b := introspection.ConvertToJson(p)
	w.Write(b)
	//fmt.Fprintf(w, renderJson(w, res))
}

func HelloHandler (w http.ResponseWriter, r * http.Request) {
    fmt.Fprintf(w, "Hey, you.")
}

