package main

import (
	"fmt"
	"net"
	"bufio"
	"io"
	"strings"
)

// var conns map[string]net.Conn
var conns = map[net.Conn]string{}

func recv(name string, conn io.Reader) {
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		message = strings.TrimPrefix(message, "\n")
		name = strings.TrimSuffix(name, "\n")
		if err != nil {
			return
		}
		fmt.Print("\n"+name+": "+string(message))
		go send(conn, message, name)
	}
}

func send(conn io.Reader, message string, name string) {
	for connection, _ := range(conns) {
		if connection != conn {
			connection.Write([]byte(name+": "+message))
		}
	}
}

func main() {
	fmt.Println("Starting server...")
	

	ln, _ := net.Listen("tcp", ":8000")
	
	defer ln.Close()
	for {
		conn, _ := ln.Accept()
		// var message string
		message, _ := bufio.NewReader(conn).ReadString('\n')
		conns[conn] = message
		go recv(string(message), conn)
	}

}