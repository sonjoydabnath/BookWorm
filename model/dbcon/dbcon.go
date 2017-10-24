package dbcon

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB
var err error

func DbConnection() {
	log.Println("Connecting to Database")
	Db, err = sql.Open("mysql", "root:mysql_root_snj@tcp(localhost:3306)/BookWorm")
	if err != nil {
		log.Println(err)
		panic(err.Error())
	} else {
		log.Println("Conneted to Database")

	}
}
