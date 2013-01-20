package router

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

const (
	CONNECT = "CONNECT"
	DELETE  = "DELETE"
	GET     = "GET"
	HEAD    = "HEAD"
	OPTIONS = "OPTIONS"
	PATCH   = "PATCH"
	POST    = "POST"
	PUT     = "PUT"
	TRACE   = "TRACE"
)

//mime-types
const (
	applicationJson = "application/json"
	applicationXml  = "applicatoin/xml"
	textXml         = "text/xml"
)

type route struct {
	method  string
	regex   *regexp.Regexp
	params  map[int]string
	handler http.HandlerFunc
}

type Router struct {
	routes  []*route
	filters []http.HandlerFunc
}

func New() *Router {
	return &Router{}
}

func (r *Router) Get(pattern string, handler http.HandlerFunc) {
	r.AddRoute(GET, pattern, handler)
}

func (r *Router) Put(pattern string, handler http.HandlerFunc) {
	r.AddRoute(PUT, pattern, handler)
}

func (r *Router) Del(pattern string, handler http.HandlerFunc) {
	r.AddRoute(DELETE, pattern, handler)
}
func (r *Router) Patch(pattern string, handler http.HandlerFunc) {
	r.AddRoute(PATCH, pattern, handler)
}

func (r *Router) Post(pattern string, handler http.HandlerFunc) {
	r.AddRoute(POST, pattern, handler)
}

func (r *Router) Static(pattern string, dir string) {
	pattern = pattern + "(.+)"
	r.AddRoute(GET, pattern, func(w http.ResponseWriter, req *http.Request) {
		path := filepath.Clean(req.URL.Path)
		path = filepath.Join(dir, path)
		http.ServeFile(w, req, path)
	})
}

func (r *Router) AddRoute(method string, pattern string, handler http.HandlerFunc) {
	parts := strings.Split(pattern, "/")
	j := 0
	params := make(map[int]string)
	for i, part := range parts {
		if strings.HasPrefix(part, ":") {
			expr := "([^/]+)"
			//a user may choose to override the defult expression
			//eg: ‘/cats/:id([0-9]+)’
			if index := strings.Index(part, "("); index != -1 {
				expr = part[index:]
				part = part[:index]
			}
			params[j] = part
			parts[i] = expr
			j++
		}
	}

	pattern = strings.Join(parts, "/")
	regex, regexErr := regexp.Compile(pattern)
	if regexErr != nil {
		// TODO avoid panic
    panic(regexErr)
		return
	}

	route := &route{}
	route.method = method
	route.regex = regex
	route.handler = handler
	route.params = params

	r.routes = append(r.routes, route)
}

func (r *Router) Filter(filter http.HandlerFunc) {
	r.filters = append(r.filters, filter)
}

func (r *Router) FilterParam(param string, filter http.HandlerFunc) {
	if !strings.HasPrefix(param, ":") {
		param = ":" + param
	}

	r.Filter(func(w http.ResponseWriter, req *http.Request) {
		p := req.URL.Query().Get(param)
		if len(p) > 0 {
			filter(w, req)
		}
	})
}

// Required by http.Handler interface. This method is invoked by the
// http server and will handle all page routing
func (r *Router) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	requestPath := req.URL.Path
	w := &responseWriter{writer: rw}

	for _, route := range r.routes {
		if req.Method != route.method || !route.regex.MatchString(requestPath) {
			continue
		}
		
		//get submatches (params)
		matches := route.regex.FindStringSubmatch(requestPath)

		//double check that the Route matches the URL pattern.
		if len(matches[0]) != len(requestPath) {
			continue
		}

		if len(route.params) > 0 {
			//add url parameters to the query param map
			values := req.URL.Query()
			for i, match := range matches[1:] {
				values.Add(route.params[i], match)
			}

			//reassemble query params and add to RawQuery
			req.URL.RawQuery = url.Values(values).Encode() + "&" + req.URL.RawQuery
			//req.URL.RawQuery = url.Values(values).Encode()
		}

		//execute middleware filters
		for _, filter := range r.filters {
			filter(w, req)
			if w.started {
				return
			}
		}
		route.handler(w, req)
		break
	}

	//no matches
	if w.started == false {
		http.NotFound(w, req)
	}
}

type responseWriter struct {
	writer  http.ResponseWriter
	started bool
	status  int
}

func (w *responseWriter) Header() http.Header {
	return w.writer.Header()
}

func (w *responseWriter) Write(p []byte) (int, error) {
	w.started = true
	return w.writer.Write(p)
}

func (w *responseWriter) WriteHeader(code int) {
	w.status = code
	w.started = true
	w.writer.WriteHeader(code)
}

func ServeJson(w http.ResponseWriter, v interface{}) {
	content, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Length", strconv.Itoa(len(content)))
	w.Header().Set("Content-Type", applicationJson)
	w.Write(content)
}

// ReadJson will parses the JSON-encoded data in the http
// Request object and stores the result in the value
// pointed to by v.
func ReadJson(req *http.Request, v interface{}) error {
	body, err := ioutil.ReadAll(req.Body)
	req.Body.Close()
	if err != nil {
		return err
	}
	return json.Unmarshal(body, v)
}

// ServeXml replies to the request with an XML
// representation of resource v.
func ServeXml(w http.ResponseWriter, v interface{}) {
	content, err := xml.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Length", strconv.Itoa(len(content)))
	w.Header().Set("Content-Type", "text/xml; charset=utf-8")
	w.Write(content)
}

// ReadXml will parses the XML-encoded data in the http
// Request object and stores the result in the value
// pointed to by v.
func ReadXml(req *http.Request, v interface{}) error {
	body, err := ioutil.ReadAll(req.Body)
	req.Body.Close()
	if err != nil {
		return err
	}
	return xml.Unmarshal(body, v)
}

// ServeFormatted replies to the request with
// a formatted representation of resource v, in the
// format requested by the client specified in the
// Accept header.
func ServeFormatted(w http.ResponseWriter, req *http.Request, v interface{}) {
	accept := req.Header.Get("Accept")
	switch accept {
	case applicationJson:
		ServeJson(w, v)
	case applicationXml, textXml:
		ServeXml(w, v)
	default:
		ServeJson(w, v)
	}

	return
}
