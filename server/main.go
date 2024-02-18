package main

import (
	"log"
	"tcpchat/server/server"
)

func main() {
	s := server.NewServer("tcp", "localhost:8000")
	log.Fatal(s.Start())
}
