package dbConfig

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

type Database struct {
}

func GetDb() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:barda87@tcp(mysql)/sample")
	if err != nil {
		panic(err)
	}
	return
}
