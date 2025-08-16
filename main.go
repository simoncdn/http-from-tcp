package main

import (
	"fmt"
	"log"
	"net"
)

const (
	port = 42069
)

func main() {
	tcpAddr := net.TCPAddr{
		Port: port,
	}

	listener, err := net.ListenTCP("tcp", &tcpAddr)
	if err != nil {
		log.Fatal("error listening tcp", err)
	}
	defer listener.Close()

	fmt.Println("Listening for TCP traffic on", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("connection error:", err)
		}

		fmt.Println("Accepted connection from", conn.RemoteAddr())
		linesChannel := getLinesChannel(conn)

		for line := range linesChannel {
			fmt.Println(line)
		}

		fmt.Println("Connection to ", conn.RemoteAddr(), "closed")
	}
}
