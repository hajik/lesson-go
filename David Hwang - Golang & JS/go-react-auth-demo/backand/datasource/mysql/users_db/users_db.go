package users_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Client   *sql.DB
	username = "root"
	password = "root"
	host     = "127.0.0.1:3306"
	schema   = "db_go_lesson"
)

func init() {

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, schema)

	var err error
	Client, err = sql.Open("mysql", dataSourceName)

	if err != nil {
		panic(err)
	}

	if err := Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("Database succesfully configured!")

}
