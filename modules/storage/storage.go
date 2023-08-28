package storage

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	connDB()

}
func GetConn() *sql.DB {
	DB, err := sql.Open("mysql", "root:Cinephile1!@tcp(127.0.0.1:3306)/cinephile?parseTime=true&charset=utf8")
	if err != nil {
		panic(err)
	}
	return DB
}
func connDB() {
	DB, err := sql.Open("mysql", "root:Cinephile1!@tcp(127.0.0.1:3306)/cinephile?parseTime=true&charset=utf8")
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
