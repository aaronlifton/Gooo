# Gooo

This is a modular web framework for the [Go language](http://www.golang.org).



## Philosophy
> "Go on Rails (Without the Rails)"

### Modular architecture
* model
  * type `map[string]interface{}`, behaves like a struct
  * implement your model interface types with anonymous BaseModel field
    * use `` `json:"-"` `` for type safety and Go future proofing
  * implement functions and variables that all models should have in Model interface
* template
  * parses templates in "tmpl/" folder and defines how they are rendered
  * currently uses the Go `"html/template"` package to parse and render
* view
  * handles routes
  * fetches rows from database as [Model](about:blank) interfaces
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

### Frameworks are dead, Long live code
* inspired by PHP, clean code architecture, and the beauty of the Go language interface type
* define your models in the model package, define your views in the view package