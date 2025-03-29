package main

import (
	"log"
	"test-auth/internal/server"
)

func main() {
	e := server.InitServer()

	log.Fatal(e.Start(":8085"))
}
