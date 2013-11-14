package model

import (
	"database/sql"
	"fmt"
	_ "github.com/bmizerany/pq"
	"github.com/aaronlifton/introspection"
	"github.com/aaronlifton/gooo/util"
	"strconv"
	"time"
	"reflect"
	"strings"
)

//dbConfig
const dbParams string = `host=ec2-XX-XXX-XXX-XXX.compute-X.amazonaws.com user=USER_NAME port=5432 password=PASS_WORD dbname=DB_NAME sslmode=require`

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

func InsertIntoDB(mod interface{}, atts util.m, atts []interface{}) {
	//db := OpenConn()
	db, err := sql.Open("postgres", dbParams)
	if err != nil {
		fmt.Println("Connection panic")
		panic(fmt.Sprintf("%s", err))
	}
	db.Begin()
	keys = strings.Join(reflect.MapKeys(atts), ",")
	values = strings.Join(introspection.MapValues(atts), ",")
	stmt, err := db.Prepare(`INSERT INTO POST (title,content,user_id,published,created,modified)
							 values ($1,$2,$3,$4,$5,$6)`)
	util.HandleErr(err)

	_, err = stmt.Exec(atts...)
	util.HandleErr(err)
	defer db.Close()
}
