package template

import (
	"bytes"
	"html/template"
	"net/http"
)

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

	/*err = tempLates.ExecuteTemplate(w, tmpl, tm)
	  if err != nil {
	      http.Error(w, err.Error(), http.StatusInternalServerError)
	  }*/
	// a developer may choose to disable template
	// caching to speedup the development process.
	// in this case we always re-parse all templates.
	if templateCache == false {
		tmp_templates := template.Must(template.ParseGlob(templatePattern))
		err = tmp_templates.ExecuteTemplate(&buf, tmpl, context)
		// in production mode, we use the cached template
		// and load the template by name
	} else {
		err = templates.ExecuteTemplate(&buf, tmpl, context)
	}
	// on error, serve internal service error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// set the content length, type, etc
	w.Write(buf.Bytes())
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}
