package storage

import (
	"cinephile/modules/env"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	connDB()
}
func GetConn() *sql.DB {
	DB, err := sql.Open("mysql", env.GetMysqlDNS())
	if err != nil {
		panic(err)
	}
	return DB
}
func connDB() {
	DB, err := sql.Open("mysql", env.GetMysqlDNS())
	if err != nil {
		panic(err)
	}
	db = DB
	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
}
func DB() *sql.DB {
	if db == nil {
		connDB()
	}
	return db
}
