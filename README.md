# Gooo

Go lang web app example showcasing PHP style web development with Go's speed. Includes batch template processing and interaction with postgresql databases.

If you like buzzwords, This is a minimal web boilerplate toolkit for the [Go language](http://www.golang.org).



## Philosophy
* Anti-magic

## Modular architecture
* model
  * `struct` type
  * embedded `Modeller` interface through `BaseModel` struct
* view
 * parses templates in `tmpl/` folder and defines how they are rendered
  * uses [`html/template`](http://golang.org/pkg/html/template/) to parse and render
  * handles routes
  * fetches rows from database as type Model interfaces
* conversion
  * models implement `interface{}` and `[]interface{}` types
  * Go's dynamic feature is interface type conversion, generally checked at runtime
  * Interface -> JSON   `interface{} -> []byte`
  * Interface -> Struct `[]interface{}` -> `map[string]interface{}`
  * 
  * GetStructValues     `interface{} -> []interface{}`
  * InterfaceName:      `interface{} -> string`
* util
  * generic error handler `HandlerErr(err error)`

## Tell me more...
### Why Go?
[Gophers](http://golang.org/doc/gopher/frontpage.png)

### Is it good?
Of course, it's written in Go. Look at that [gopher](http://golang.org/doc/gopher/frontpage.png).

## Let's Gooo test the example blog app
* resolve dependencies
  * go get https://github.com/bmizerany/pq (pure Go postgres driver for `database/sql`)
  * don't want to use postgresql? Use it anyway.
    * Sign up for a free Heroku Postgres account [here](https://postgres.heroku.com/).
    * Create a database and save the connection params for the next step.
* configure the database connection variable `dpParams` in the model package (`model/model.go`)
* `go install` iff first build
* `go build && ./gooo`
* [http://localhost:8080](http://localhost:8080)
* Gooo celebrate
* Gooo outside

## Let's Gooo write your own Gooo app
* resolve dependencies
  * go get https://github.com/bmizerany/pq (pure Go postgres driver for `database/sql`)
  * don't want to use postgresql? Use it anyway.
    * Sign up for a free Heroku Postgres account [here](https://postgres.heroku.com/).
    * Create a database and save the connection params for the next step.
* define your model interfaces and configure the database connection in the model package (`model/model.go`)
  * implement your model interface types with anonymous `BaseModel` field
    * use `` `json:"-"` `` for type safety and Go lang future proofing
  * implement functions and variables available to all models with anonymous `BaseModel` field in the `Model` interface type
* define your views as request handler functions in the view package (`view/view.go`)
* write your templates in `tmpl/` ([Go text/template syntax](http://golang.org/pkg/text/template/))
* define routes in main package gooo.go
* `go install` iff first build
* `go build && ./gooo`
* [http://localhost:8080](http://localhost:8080)
* Gooo celebrate
* Gooo outside

## Gooo read these
* [text/template](http://golang.org/pkg/text/template/)
* [html/template](http://golang.org/pkg/html/template/)
* [net/http](http://golang.org/pkg/net/http/)
* [database/sql](http://golang.org/pkg/database/sql/)
* [encoding/json](http://golang.org/pkg/encoding/json/)

- - -

Enjoy,

  \- Aaron Lifton