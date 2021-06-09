package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatalf("Error listening to TCP: %v\n", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error acceptinf connection: %v\n", err)
			continue
		}

		// If go-routine not used here, then only one client can connect.
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, "response from server\n")
		if err != nil {
			return
		}
		time.Sleep(time.Second)
	}
}
