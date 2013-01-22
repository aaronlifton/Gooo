package model

import (
	"database/sql"
	"fmt"
	_ "github.com/bmizerany/pq"
	"gooo/util"
	"strconv"
	"time"
)

//dbConfig
//const dbParams string = `host=ec2-XX-XXX-XXX-XXX.compute-X.amazonaws.com user=USER_NAME port=5432 password=PASS_WORD dbname=DB_NAME sslmode=require`
const dbParams string = `host=ec2-54-243-239-221.compute-1.amazonaws.com user=kwcwqwdgfelhrs port=5432 password=KKKZ3FadJRB_0IC8PK32KoKpti dbname=d62335du1mgsdc sslmode=require`

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

func TestEmptyDB() bool {
	db, err := sql.Open("postgres", dbParams)
	if err != nil {
		fmt.Println("Connection panic")
		panic(fmt.Sprintf("%s", err))
	}
	db.Begin()
	q := `select relname from pg_class where relname = 'post' and relkind='r'`
	var initialized bool = false
	rows, err := db.Query(q)
	if err != nil {
		fmt.Println("DB Query panic")
		panic(fmt.Sprintf("%s", err))
	}
	for rows.Next() {
		var relname string
		err = rows.Scan(&relname)
		if len(relname) > 0 {
			initialized = true
			return initialized
		}
	}
	err = rows.Err()
	util.HandleErr(err)
	rows.Close()

	if initialized == false {
		fmt.Println("\033[32;1m %s\033[0m", "Initializing empty DB")
		q = `DROP TABLE IF EXISTS post CASCADE;
	       CREATE TABLE post(id INTEGER PRIMARY KEY NOT NULL DEFAULT nextval('post_id_seq'),
	                title VARCHAR(32), content TEXT, user_id INTEGER,
	                published BOOLEAN, created TIMESTAMP, modified TIMESTAMP)`
		_, err = db.Exec(q)
		util.HandleErr(err)
	}
	defer db.Close()
	return initialized
}

func OpenConn() *sql.DB {
	db, err := sql.Open("postgres", dbParams)
	if err != nil {
		fmt.Println("Connection panic")
		panic(fmt.Sprintf("%s", err))
	}
	db.Begin()
	return db
}

func InsertIntoDB(atts []interface{}) {
	//db := OpenConn()
	db, err := sql.Open("postgres", dbParams)
	if err != nil {
		fmt.Println("Connection panic")
		panic(fmt.Sprintf("%s", err))
	}
	db.Begin()
	stmt, err := db.Prepare(`INSERT INTO POST (title,content,user_id,published,created,modified)
							 values ($1,$2,$3,$4,$5,$6)`)
	util.HandleErr(err)

	_, err = stmt.Exec(atts...)
	util.HandleErr(err)
	defer db.Close()
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
