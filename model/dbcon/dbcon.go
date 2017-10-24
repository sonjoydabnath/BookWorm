package dbcon

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sonjoydabnath/BookWorm/model/configs"
)

var Db *sql.DB
var err error

func DbConnection(config configs.Config) {
	log.Println("Connecting to Database at " + config.Database.Host + ":" + config.Database.Port)
	Db, err = sql.Open("mysql", config.Database.Username+":"+config.Database.Password+"@tcp("+config.Database.Host+":"+config.Database.Port+")/"+config.Database.Schema)
	if err != nil {
		log.Println(err)
		panic(err.Error())
	} else {
		log.Println("Conneted to Database")

	}
}
