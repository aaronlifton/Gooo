package model

import (
	"database/sql"
	"fmt"
	_ "github.com/bmizerany/pq"
	"goweb/util"
	"time"
)

const dbParams string = `host=ec2-XX-XXX-XXX-XXX.compute-X.amazonaws.com user=USER_NAME port=5432 password=PASS_WORD dbname=DB_NAME sslmode=require`

var (
	M BaseModel
)

type Model interface {
	FindById()
	FindAll()
}
type BaseModel struct {
	Model
	Name string
}

type Post struct {
	BaseModel `json:"-"`
	Id        int
	Title     string
	Content   string //[]byte
	UserId    int
	Published bool
	Created   time.Time
	Modified  time.Time
}

func TestEmptyDB(db *sql.DB) {
	testQ := `select relname from pg_class where relname = 'post' and relkind='r'`
	var initialized int = 0
	rows, err := db.Query(testQ)
	if err != nil {
		fmt.Println("DB Query panic")
		panic(fmt.Sprintf("%s", err))
	}
	for rows.Next() {
		var relname string
		err = rows.Scan(&relname)
		if len(relname) > 0 {
			initialized++
		}
	}
	err = rows.Err()
	util.HandleErr(err)
	rows.Close()

	if initialized == 0 {
		q := `DROP SEQUENCE IF EXISTS post_id_seq CASCADE;
	            DROP TABLE IF EXISTS post CASCADE;
	            CREATE SEQUENCE post_id_seq;
	            CREATE TABLE post(id INTEGER PRIMARY KEY NOT NULL DEFAULT nextval('post_id_seq'),
	                title VARCHAR(32), content TEXT, user_id INTEGER,
	                published BOOLEAN, created TIMESTAMP, modified TIMESTAMP)`
		_, err = db.Exec(q)
		util.HandleErr(err)
	}
}

func OpenConn() *sql.DB {
	db, err := sql.Open("postgres", dbParams)
	if err != nil {
		fmt.Println("Connection panic")
		panic(fmt.Sprintf("%s", err))
	}
	db.Begin()
	//defer db.Close()
	return db
}

func InsertIntoDB(atts []interface{}) {
	db := OpenConn()
	stmt, err := db.Prepare(`INSERT INTO POST (title,content,user_id,published,created,modified)
							 values ($1,$2,$3,$4,$5,$6)`)
	util.HandleErr(err)

	_, err = stmt.Exec(atts...)
}
