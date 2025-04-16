package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

const numClients = 100

func startClient(addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatalf("Error connecting to %v: %v", addr, err)
	}

	localAddr := conn.LocalAddr()

	defer func() {
		fmt.Printf("Client %v disconnected\n", localAddr)
		conn.Close()
	}()

	for {
		_, err := conn.Write([]byte("ping"))
		if err != nil {
			log.Fatalf("Error writing to %v: %v", localAddr, err)
		}

		time.Sleep(1 * time.Second)
	}
}

func main() {
	for range numClients {
		go startClient("localhost:8080")
	}

	time.Sleep(1 * time.Minute)
}
