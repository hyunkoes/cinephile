package main

import (
	"cinephile/modules/server"
)

var mode = 0

func main() {
	server.Serve(mode)
}
