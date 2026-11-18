package main

import (
	"fmt"
	"strings"
	server "github.com/LanceC93/sserver/server"
	client "github.com/LanceC93/sserver/client"
)

type Server interface {
	StartServer()
}

type Client interface {
	StartClient()
}

func main() {
	
	fmt.Printf("(s)erver or (c)lient: ")
	var input string
	fmt.Scan(&input)

	switch strings.ToLower(input) {
		case "server", "s":
			var s Server = server.Chat{}
			s.StartServer()
		case "client", "c":
			var c Client = client.Chat{}
			c.StartClient()
		default:
			fmt.Printf("Invalid input.")
	}
}

