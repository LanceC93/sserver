package main

import (
	"fmt"
	"strings"
	chat "github.com/LanceC93/sserver/chat"
	"os"
)

type Server interface {
	StartServer()
	StartClient()
}

var serverTypes = map[string]Server {
	"chat": chat.Chat{},
	"proxy": nil,
}

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("No arguments passed.")
		return
	}

	serveType := os.Args[1]

	if _, exists := serverTypes[serveType]; !exists {
		fmt.Printf("Invalid Argument.")
		return
	}


	fmt.Printf("(s)erver or (c)lient: ")
	var input string
	fmt.Scan(&input)
	var s Server = serverTypes[serveType]

	switch strings.ToLower(input) {
		case "server", "s":
			s.StartServer()
		case "client", "c":
			s.StartClient()
		default:
			fmt.Printf("Invalid input.")
	}
}

