package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const (
	network = "udp"
	address = "localhost:42069"
)

func main() {
	rAddr, err := net.ResolveUDPAddr(network, address)
	if err != nil {
		log.Fatal("couldn't resolve udp address:", err)
	}

	conn, err := net.DialUDP(network, nil, rAddr)
	if err != nil {
		log.Fatal("couldn't open udp connection:", err)
	}
	defer conn.Close()

	fmt.Printf("Sending to %s. Type your message and press Enter to send. Press Ctrl+C to exit.\n", address)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")

		msg, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("couldn't read string", err)
		}

		_, err = conn.Write([]byte(msg))
		if err != nil {
			log.Fatal("connection writing error:", err)
		}

		fmt.Printf("Message sent: %s", msg)
	}
}
