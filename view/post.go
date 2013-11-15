package view

import (
  "github.com/aaronlifton/gooo/introspection"
  _ "github.com/aaronlifton/gooo/memory"
  "github.com/aaronlifton/gooo/model"
  "net/http"
  "strconv"
  "time"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
  db := model.OpenConn()
  //TODO: cache this result
  //model.TestEmptyDB()
  latestPosts := model.GetPosts(10)
  ctx := m{"latestPosts": latestPosts}
  defer db.Close()
  //var listTmpl = template.Must(template.ParseFiles("tmpl/base.html","tmpl/index.html"))
  //listTmpl.ExecuteTemplate(w,"index", ctx)
  //listTmpl.ExecuteTemplate(w,"base",  nil)
  //listTmpl.Execute(w, nil)
  RenderTemplate(w, "base", ctx)
}

func NewPostHandler(w http.ResponseWriter, r *http.Request) {
  title := r.FormValue("title")
  body := r.FormValue("content")
  userId, err := strconv.Atoi(r.FormValue("userId"))
  if err != nil {
    userId = 0
    //TODO: implement user model
  }
  published := true
  p := model.Post{0, title, body, userId, published, time.Now(), time.Now()}
  atts := introspection.GetStructValues(&p)
  model.InsertIntoDB(atts)
  http.Redirect(w, r, "/", http.StatusFound)
}

func PostJSONHandler(w http.ResponseWriter, r *http.Request) {
  var p model.Post = model.Post{0, "Test", "test post please ignore", 1, true, time.Now(), time.Now()}
  w.Header().Set("Content-Type", "application/json")
  b := introspection.ConvertToJson(p)
  w.Write(b)
}