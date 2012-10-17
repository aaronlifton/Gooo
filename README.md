# Gooo

This is a modular web framework for the [Go language](http://www.golang.org).



## Philosophy
> "Go on Rails (Without the Rails)"

> What is a framework? framework (noun): an essential supporting structure of a building, vehicle, or object, a basic structure underlying a system, concept, or text

* Inspired by PHP, clean code architecture, and the beauty of the Go language and its interface type

* Anti-Magic

## Modular architecture
* model
  * type `map[string]interface{}`, behaves like a struct
  * implement your model interface types with anonymous `BaseModel` field
    * use `` `json:"-"` `` for type safety and Go lang future proofing
  * implement functions and variables available to all models with anonymous `BaseModel` field in the `Model` interface type
* template
  * parses templates in `tmpl/` folder and defines how they are rendered
  * currently uses the Go [`html/template`](http://golang.org/pkg/html/template/) package to parse and render
* view
  * handles routes
  * fetches rows from database as type Model interfaces
* conversion
  * everything in Go descends from the interface type (essentially)
  * to ease interaction with templates and databases, they must be converted
    to and from the `interface{}` and `[]interface{}` types
  * Interface -> JSON
  * Interface -> Struct
  * `[]interface{}` -> `map[string]interface{}`
  * GetStructValues
  * StructName: gets name of a []interface{} type
* util
  * generic error handler `HandlerErr(err error)`

## Tell me more...
### Is this really a framework? Why isn't it magic?
Yes.

### Is it good?
Yes.

## Let's Gooo
* resolve dependencies
  * go get https://github.com/bmizerany/pq (pure Go postgres driver for database/sql)
  * don't want to use postgresql? Use it anyway.
    * Sign up for a free Heroku Postgres dev account [here](https://postgres.heroku.com/).
    * Create a database and save the connection params for the next step.
* configure the database connection variable `dpParams` in the model package (`model/model.go`)
* `go install` iff first build
* `go build && ./gooo`
* [http://localhost:8080](http://localhost:8080)
* Gooo celebrate

## Let's Gooo write your own Gooo app
* resolve dependencies
  * go get https://github.com/bmizerany/pq (pure Go postgres driver for database/sql)
  * don't want to use postgresql? Use it anyway.
    * Sign up for a free Heroku Postgres dev account [here](https://postgres.heroku.com/).
    * Create a database and save the connection params for the next step.
* define your model interfaces and configure the database connection in the model package (`model/model.go`)
* define your views as request handler functions in the view package
* write your templates in `tmpl/` ([Go text/template syntax](http://golang.org/pkg/text/template/))
* define routes in main package gooo.go
* `go install` iff first build
* `go build && ./gooo`
* [http://localhost:8080](http://localhost:8080)
* Gooo celebrate

## Read these before use
* [text/template](http://golang.org/pkg/text/template/)
* [html/template](http://golang.org/pkg/html/template/)
* [net/http](http://golang.org/pkg/net/http/)
* [database/sql](http://golang.org/pkg/database/sql/)
* [encoding/json](http://golang.org/pkg/encoding/json/)

- - -

Enjoy,

  \- Aaron Lifton