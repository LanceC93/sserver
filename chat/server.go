package chat

import(
	"fmt"
	"net"
	"log"
	"bufio"
	"errors"
	"io"
	"strings"
)

const (
	SERVER_HOST="localhost"
	SERVER_PORT="9988"
	SERVER_TYPE="tcp"
)

type Chat struct {}
//creates a server
func (Chat) StartServer() {
	//initialize connection
	ln, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		log.Fatalf("Failed to start server: %w", err)
	}
	
	//listen for connections
	fmt.Printf("Listening...\n")
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("Failed to accept conection: %w", err)
		}
		
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Printf("Connection established.\n")

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	exit := false

	for !exit {
		input, err := reader.ReadString('\n')

		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Printf("Connection to client closed.\n")
			} else {
				fmt.Printf("Problem reading from client: %w", err)
			}
			exit = true

		}
			fmt.Printf("Received.\n")
			writer.WriteString(strings.ToUpper(input))
			writer.Flush()
	}

	fmt.Printf("Closing...")
	conn.Close()
}
