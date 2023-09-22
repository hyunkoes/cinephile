package main

import (
	"cinephile/modules/env"
	"cinephile/modules/server"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var mode = 0

func main() {
	if os.Getenv(`env`) == "" {
		godotenv.Load(`.env.local`)
		fmt.Println(env.GetMysqlDNS())
	}
	server.Serve(mode)
}
