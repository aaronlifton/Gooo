# Gooo

```
   _____
  / ____|
 | |  __  ___   ___   ___
 | | |_ |/ _ \ / _ \ / _ \
 | |__| | (_) | (_) | (_) |
  \_____|\___/ \___/ \___/
```

Pronunciation: */ɡoʊʊʊ/*

Go lang web app "framework" showcasing straightforward, no-magic, web development with the [Go language](http://www.golang.org). Includes batch template processing and interaction with postgresql databases, and Model-View architecture.


## Philosophy
* Anti-magic
* So simple, that it's complex.
* So complex, that it works.
* If it doesn't work, publish it.

## Modular architecture
* model
  * `struct` type
  * no special tags or fields, models are just Go structs
  * business logic - straightforward DB methods to be used in the view module
* view
 * parses templates in `tmpl/` folder and defines how they are rendered
  * uses [`html/template`](http://golang.org/pkg/html/template/) to parse and render
  * fetches rows from database as type Model interfaces
* router
  * handles dynamic and static routes, request methods, main handler matches the request URL against the routes
  * middleware filters for restful routes
* introspection
  * models implement `interface{}` and `[]interface{}` types
  * Go's dynamic feature is interface type conversion, generally checked at runtime
  * Interface -> JSON   `interface{} -> []byte`
  * Interface -> Struct `[]interface{}` -> `map[string]interface{}`
  *
  * GetStructValues     `interface{} -> []interface{}`
  * InterfaceName:      `interface{} -> string`
* util
  * generic error handler `HandlerErr(err error)`

---

![Martin Odersky](http://i.imgur.com/jB8aa.jpg?1)

*Tested and approved by Typeunsafe&copy; Corporation*

---

## Tell me more...
### Why Go?
* [Gophers](http://golang.org/doc/gopher/frontpage.png)
* Hip
* New
* Unproven

### Is it good?
Are you good?

## Let's Gooo test the example blog app
* resolve dependencies and install
  * `./install`
  * don't want to use postgresql? Use it anyway.
    * Sign up for a free Heroku Postgres account [here](https://postgres.heroku.com/).
    * Create a database and save the connection params for the next step.
* configure the database connection variable `dbParams` in the model package (`model/model.go`)
* `./gooo`
* [http://localhost:8080](http://localhost:8080)
* Gooo celebrate
* Gooo outside

## Let's Gooo write your own Gooo app
* resolve dependencies and install
  * `./install`
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
* `./gooo`
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

<p align="center">
  <img src="http://i.imgur.com/NSscm.jpg" alt="Gopher"/>
</p>


- - -

Enjoy,

  \- Aaron Lifton
