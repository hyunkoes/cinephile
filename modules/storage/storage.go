package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	connDB()

}
func GetConn() *sql.DB {
	DB, err := sql.Open("mysql", "root:Cinephile1!@tcp(127.0.0.1:3306)/cinephile?parseTime=true")
	if err != nil {
		panic(err)
	}
	return DB
}
func connDB() {
	fmt.Println("CONN DB FUNC")
	DB, err := sql.Open("mysql", "root:Cinephile1!@tcp(127.0.0.1:3306)/cinephile?parseTime=true")
	if err != nil {
		panic(err)
	}
	db = DB
	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("! Connected to:", version)
	fmt.Println("CONN! PING TEST : ", db.Ping())
}
func DB() *sql.DB {
	fmt.Println("DB() Called!")
	if db == nil {
		fmt.Println("DB IS NULL")
		connDB()
	}
	return db
}
