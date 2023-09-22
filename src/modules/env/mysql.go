package env

import (
	"fmt"
	"os"
)

func GetMysqlDNS() string {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8", os.Getenv(`MYSQL_USER`), os.Getenv(`ROOT_PASSWORD`), os.Getenv(`MYSQL_HOST`), os.Getenv(`MYSQL_PORT`), os.Getenv(`MYSQL_DATABASE`))
	return dns
}
