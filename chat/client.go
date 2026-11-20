package chat

import (
	"fmt"
	"net"
	"bufio"
	"strings"
	"errors"
	"io"
	"sync/atomic"
	"os"
)

//create client
func (Chat) StartClient() {
	//start connection
	conn, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		fmt.Errorf("Unable to connect to %s:%s: %s", SERVER_HOST, SERVER_PORT, err)
	}
	fmt.Printf("Connection established. (q)uit to exit\n")
	
	var exit atomic.Bool
	exit.Store(false)

	go clientRead(conn, &exit)
	clientWrite(conn, &exit)

	fmt.Printf("Closing...")
	conn.Close()
}

//process for reading
func clientRead(conn net.Conn, exit *atomic.Bool) {

	reader := bufio.NewReader(conn)

	for !exit.Load() {
		response, err := reader.ReadString('\n')

		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Printf("Connection to server closed.\n")
				exit.Store(true)
			} else {
				fmt.Printf("Problem reading from server: %w", err)
			}
			exit.Store(true)
		} else {
			fmt.Printf("server > %sclient > ", response)
		}
	}
}

//process for writing
func clientWrite(conn net.Conn, exit *atomic.Bool) {
	var input string
	in := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(conn)
	fmt.Printf("client > ")

	//client sends to the server
	for !exit.Load() {
		input,_ = in.ReadString('\n')

		if strings.ToLower(input) == "quit\n" || strings.ToLower(input) == "q\n" {
			exit.Store(true)
			conn.Close()
		} else {
			writer.WriteString(input)
			writer.Flush()
		}
	}
}
