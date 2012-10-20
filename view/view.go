package view

import (
	"fmt"
	"gooo/conversion"
	"gooo/model"
	tmpl "gooo/template"
	"html/template"
	"net/http"
	"time"
)

var (
	templateCache          = false
	templatePattern string = "tmpl/*.html"
	templates       *template.Template
)

func HomeHandler(w http.ResponseWriter, r *http.Request, title string) {
	db := model.OpenConn()
	model.TestEmptyDB(db)
	var p model.Post = model.Post{model.M, 2, "Test", "test post please ignore", 1, true, time.Now(), time.Now()}
	var p2 model.Post = model.Post{model.M, 2, "Test2", "another test post please ignore", 1, true, time.Now(), time.Now()}
	fmt.Println(p.ModelName())
	/*atts := conversion.GetStructValues(&p)
	for z := range atts {
		fmt.Println(atts[z])
	}*/
	/*tmt, err := db.Prepare(`INSERT INTO POST (title,content,user_id,published,created,modified)
							 values ($1,$2,$3,$4,$5,$6)`)
	HandleErr(err)

	_, err = stmt.Exec(atts...)
	*/
	//post := conversion.ConvertToJson(p)
	/*var f interface{}
	var b []byte
	b, err = json.Marshal(p)
	HandleErr(err)
	err = json.Unmarshal(b, &f)
	m := f.(map[string]interface{})*/
	x := map[string]interface{}{"p1": conversion.ConvertToMap(p), "p2": conversion.ConvertToMap(p2)}
	y := map[string]interface{}{"posts": x}
	//util.HandleErr(err)
	defer db.Close()
	//defer stmt.Close()
	//osts := p.interfaceify()
	tmpl.RenderTemplate(w, "index", y)
}

///*func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
//	p := &Post{2, "Wu Tang Clan", "dude whats up lol", 1, true, time.Now(), time.Now()}
//	/*
//	   if err != nil {
//	       http.Redirect(w, r, "/edit/"+title, http.StatusFound)
//	       return
//	   }*/
//	pMap := p.interfaceify()
//	renderTemplate(w, "view", pMap)
//}
//
//func editHandler(w http.ResponseWriter, r *http.Request, title string) {
//	p := &Post{2, "Wu Tang Clan", "dude whats up lol", 1, true, time.Now(), time.Now()}
//	/*if err != nil {
//	    p = &Post{Title: title}
//	}*/
//	pMap := p.interfaceify()
//	renderTemplate(w, "edit", pMap)
//}
//*/

func SaveHandler(w http.ResponseWriter, r *http.Request, title string) {
	content := r.FormValue("content")
	p := &model.Post{Title: title, Content: content, UserId: 1, Published: true, Created: time.Now(), Modified: time.Now()}
	p.Title = "lol"
	/*err := p.save()
	  if err != nil {
	      http.Error(w, err.Error(), http.StatusInternalServerError)
	      return
	  }*/
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
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
	b := conversion.ConvertToJson(p)
	w.Write(b)
	//fmt.Fprintf(w, renderJson(w, res))
}
