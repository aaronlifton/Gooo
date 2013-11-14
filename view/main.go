package view

func HelloHandler(w http.ResponseWriter, r *http.Request) {
  params := r.URL.Query()
  lastName := params.Get(":last")
  firstName := params.Get(":first")
  name := m{"content": m{"firstName": firstName, "lastName": lastName}}
  RenderTemplate(w, "base", name)
}

func CountHandler(w http.ResponseWriter, r *http.Request) {
  sess := globalSessions.SessionStart(w, r)
  createtime := sess.Get("createtime")
  if createtime == nil {
    sess.Set("createtime", time.Now().Unix())
  } else if (createtime.(int64) + 360) < (time.Now().Unix()) {
    globalSessions.SessionDestroy(w, r)
    sess = globalSessions.SessionStart(w, r)
  }
  ct := sess.Get("countnum")
  if ct == nil {
    sess.Set("countnum", 1)
  } else {
    sess.Set("countnum", (ct.(int) + 1))
  }
  t, _ := template.ParseFiles("count.html")
  w.Header().Set("Content-Type", "text/html")
  t.Execute(w, sess.Get("countnum"))
}