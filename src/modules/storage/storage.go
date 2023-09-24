package storage

import (
	"cinephile/modules/env"
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func init() {
	if os.Getenv(`env`) == "" {
		godotenv.Load(`.env.local`)
	}
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
