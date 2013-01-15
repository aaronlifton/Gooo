package model

import (
	"database/sql"
	"fmt"
	_ "github.com/bmizerany/pq"
	"gooo/introspection"
	"gooo/util"
	"time"
)

//dbConfig
const dbParams string = `host=ec2-XX-XXX-XXX-XXX.compute-X.amazonaws.com user=USER_NAME port=5432 password=PASS_WORD dbname=DB_NAME sslmode=require`

var (
	M BaseModel
)

type Modeller interface {
	FindById() int
	FindAll() int
	ModelName() string
}

type BaseModel struct {
	Modeller `json:"-"`
	//Name  string
}

// example model
// note anonymous field BaseModel
// and its json tag
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

/*func (p Post) ModelName() string {
	return introspection.InterfaceName(p)
}*/
func (b BaseModel) ModelName() string {
	return introspection.InterfaceName(b)
}

func TestEmptyDB(db *sql.DB) {
	q := `select relname from pg_class where relname = 'post' and relkind='r'`
	var initialized int = 0
	rows, err := db.Query(q)
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
		q = `DROP SEQUENCE IF EXISTS post_id_seq CASCADE;
	            DROP TABLE IF EXISTS post CASCADE;
	            CREATE SEQUENCE post_id_seq;
	            CREATE TABLE post(id INTEGER PRIMARY KEY NOT NULL DEFAULT nextval('post_id_seq'),
	                title VARCHAR(32), content TEXT, user_id INTEGER,
	                published BOOLEAN, created TIMESTAMP, modified TIMESTAMP)`
		_, err = db.Exec(q)
		util.HandleErr(err)
	}
	defer db.Close()
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

func GetPosts(n int) {
	db := OpenConn()
	rows, err := db.Query(`SELECT * FROM POST LIMIT 10`)
	util.HandleErr(err)
	defer db.Close()
	rows.Next()
	cols, _ := rows.Columns()
	out := make([]interface{}, len(cols))
	dest := make([]interface{}, len(cols))
	for i, _ := range dest {
		dest[i] = &out[i]
	}
	err = rows.Scan(dest...)

	fmt.Sprintf("%s", dest)
	fmt.Println(rows)
	//fmt.Println((*valuePtrs[0].(*interface{})).(string))
	//return rows
}
