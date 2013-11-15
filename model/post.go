package model

import(
  "strconv"
  "time"
)

// example model
type Post struct {
  Id        int
  Title     string
  Content   string
  UserId    int
  Published bool
  Created   time.Time
  Modified  time.Time
}

func GetPosts(n int) (results []interface{}) {
  db := OpenConn()
  stmt, err := db.Prepare(`SELECT * from post order by created DESC LIMIT $1`)
  util.HandleErr(err)

  var atts = make([]interface{}, 0)
  atts = append(atts, strconv.Itoa(n))
  rows, err := stmt.Query(atts...)
  util.HandleErr(err)

  defer db.Close()
  //fields, err := rows.Columns()
  util.HandleErr(err)

  results = make([]interface{}, 0)

  var id, userId int
  var title string
  var content string
  var published bool
  var created, modified time.Time

  for rows.Next() {
    err = rows.Scan(&id, &title, &content, &userId, &published, &created, &modified)
    util.HandleErr(err)
    var p = Post{id, title, content, userId, published, created, modified}
    results = append(results, p)
  }
  return results
}
