package main

import (
	"fmt"
	chat "github.com/LanceC93/sserver/chat"
	"os"
	"flag"
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
	isServer := flag.Bool("s", false, "is this a server.")
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Printf("No arguments passed.")
		return
	}

	serveType := flag.Arg(0)

	if _, exists := serverTypes[serveType]; !exists {
		fmt.Printf("Invalid Argument.")
		return
	}

	var s Server = serverTypes[serveType]

	if *isServer {
		s.StartServer()
	} else {
		s.StartClient()
	}
}

