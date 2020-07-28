package main

import (
	"fmt"
	"net"
	"bufio"
	"io"
	"os"
	"strings"
)

func recv(conn io.Reader, messages chan string) {
	// tmp := make([]byte, 4096)
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		message = strings.TrimSuffix(message, "\n")
		messages<-message
	}
}


func send(sends chan string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		sends<-text
	}
}


func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8000")
	// Sends username to server
	name := os.Args[1]
	fmt.Fprintf(conn, name + "\n")

	messages := make(chan string)
	sends := make(chan string)
	go recv(conn, messages)
	go send(sends)
	for {
		select {
		case message := <- messages:
			fmt.Println(string(message))
		case text := <- sends:
			fmt.Fprintf(conn, text + "\n")
		}
	}
}