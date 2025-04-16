package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	defer func() {
		fmt.Println("Connection closed:", conn.RemoteAddr())
		conn.Close()
	}()
	_, err := io.Copy(io.Discard, conn)
	if err != nil {
		log.Fatalf("Error reading from %v: %v", conn.RemoteAddr(), err)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Server start error:", err)
	}

	defer listener.Close()
	fmt.Println("Server started on:", listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Temporary() {
				fmt.Println("Temporary error:", err)
				continue
			}
			log.Fatal("Critical Accept error", err)
		}
		go handleConnection(conn)
	}
}
