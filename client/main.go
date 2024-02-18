package main

import (
	"tcpchat/client/client"
)

func main() {
	c := client.NewClient("tcp", "localhost:8000")

	c.Start()
}
