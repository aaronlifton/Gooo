package view

import (
	"bytes"
	"gooo/introspection"
	"gooo/model"
	"html/template"
	"net/http"
	"strconv"
	"time"
  "fmt"
  "crypto/hmac"
  "crypto/sha1"
  "encoding/base64"
  "strings"
  "gooo/session"
  _"gooo/memory"
)

// template config
var (
	templateCache   bool   = false
	templatePattern string = "tmpl/*.html"
	templates       *template.Template
  CookieSecret =  "7C19QRmwf3mHZ9CPAaPQ0hsWeufKd"
  globalSessions  *session.Manager
)

type m map[string]interface{}

type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

func ParseTemplateGlob(pattern string, cache bool) {
	templateCache = cache
	templatePattern = pattern
	templates = template.Must(template.ParseGlob(templatePattern))
}

func RenderTemplate(w http.ResponseWriter, tmpl string, context m) {
	var err error
	var buf bytes.Buffer

	// you can disable template caching for speed
	// otherwise we always re-parse the templates
	if templateCache == false {
		tmp_templates := template.Must(template.ParseGlob(templatePattern))
		err = tmp_templates.ExecuteTemplate(&buf, tmpl, context)
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

func TestCookieSetHandler(w http.ResponseWriter, r *http.Request) {
	//var p model.Post = model.Post{0, "Suave! Goddamn you're one suave fucker!", "You stay alive, baby. Do it for Van Gogh.", 1, true, time.Now(), time.Now()}
	//var p2 model.Post = model.Post{0, "Test Post", "Heineken? Fuck that shit! Pabst Blue Ribbon!", 1, true, time.Now(), time.Now()}
	//posts := m{"p1": introspection.ConvertToMap(p), "p2": introspection.ConvertToMap(p2)}
  encodedValue := base64.URLEncoding.EncodeToString([]byte(CookieSecret))
  SetSecureCookie(w,r,"user",encodedValue,3600)
  SetCookie(w,r,"visited","true",0)
  ctx := m{"content": "cookie set."}
	RenderTemplate(w, "base", ctx)
}
func TestCookieGetHandler(w http.ResponseWriter, r *http.Request) {
  encodedVal, _ := GetSecureCookie(w, r, "user")
  unsecureCookieVal, _ := GetCookie(w, r, "visited")
  ctx := m{"content": "Your secure cookie is: " + encodedVal + "Your insecure cookie is: " + unsecureCookieVal}
  RenderTemplate(w, "base", ctx)
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
    } else if (createtime.(int64) + 360) < (time.Now (). Unix ()) {
        globalSessions.SessionDestroy(w, r)
        sess = globalSessions.SessionStart(w, r)
    }
    ct := sess.Get("countnum")
    if ct == nil {
        sess.Set("countnum", 1)
    } else {
        sess.Set("countnum", (ct.(int) + 1))
    }
    t, _ := template.ParseFiles ("count.html")
    w.Header().Set ("Content-Type", "text/html")
    t.Execute(w, sess.Get("countnum"))
}

//Sets a cookie -- duration is the amount of time in seconds. 0 = forever
func SetCookie(w http.ResponseWriter, r *http.Request, name string, value string, age int64) {
    var utctime time.Time
    if age == 0 {
        // 2^31 - 1 seconds (roughly 2038)
        utctime = time.Unix(2147483647, 0)
    } else {
        utctime = time.Unix(time.Now().Unix()+age, 0)
    }
    fmt.Println(utctime)
    timeStr := fmt.Sprintf("%d", time.Now().Unix())

    cookie := fmt.Sprintf("%s=%s; expires=%s", name, value, timeStr)
    w.Header().Set("Set-Cookie", cookie)
}

func SetSecureCookie(w http.ResponseWriter, r *http.Request, name string, val string, age int64) {
    cVal := signedCookieValue(CookieSecret, name, val)
    SetCookie(w, r, name, cVal, age)
  }

func GetSecureCookie(w http.ResponseWriter, r *http.Request, name string) (string, bool) {
    for _, cookie := range r.Cookies() {
        if cookie.Name != name {
            continue
        }
        return validateCookie(cookie, CookieSecret)
    }
    return "", false
}

func GetCookie(w http.ResponseWriter, r *http.Request, name string) (string, bool) {
    for _, cookie := range r.Cookies() {
        if cookie.Name != name {
            continue
        }
        return string(cookie.Value), true
    }
    return "", false
}

func validateCookie(cookie *http.Cookie, seed string) (string, bool) {
	// value, timestamp, sig
	parts := strings.Split(cookie.Value, "|")
	if len(parts) != 3 {
		return "", false
	}
  ts,  _ := strconv.ParseInt(parts[1],0,64)
  if time.Now().Unix()-31*86400 > ts {
    return "", false
  }
	sig := cookieSignature(seed, cookie.Name, parts[0], parts[1])
	if parts[2] == sig {
		//ts, err := strconv.Atoi(parts[1])
		//if err == nil {// && int64(ts) > time.Now().Add(time.Duration(24)*7*time.Hour*-1).Unix() {
			// it's a valid cookie. now get the contents
			rawValue, err := base64.URLEncoding.DecodeString(parts[0])
			if err == nil {
				return string(rawValue), true
			}
		//}
	}
	return "", false
}

func signedCookieValue(seed string, key string, value string) string {
	encodedValue := base64.URLEncoding.EncodeToString([]byte(value))
	timeStr := fmt.Sprintf("%d", time.Now().Unix())
	sig := cookieSignature(seed, key, encodedValue, timeStr)
	cookieVal := fmt.Sprintf("%s|%s|%s", encodedValue, timeStr, sig)
	return cookieVal
}

func cookieSignature(args ...string) string {
	h := hmac.New(sha1.New, []byte(args[0]))
	for _, arg := range args[1:] {
		h.Write([]byte(arg))
	}
	var b []byte
	b = h.Sum(b)
	return base64.URLEncoding.EncodeToString(b)
}

