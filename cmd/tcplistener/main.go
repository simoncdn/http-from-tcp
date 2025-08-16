package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
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

func getLinesChannel(f io.ReadCloser) <-chan string {
	ch := make(chan string)

	go func(){
		defer f.Close()
		defer close(ch)

		currentLine := ""

		for {
			buffer := make([]byte, 8)

			n, err := f.Read(buffer)
			if err != nil {
				if currentLine != "" {
					ch <- currentLine
				}

				if errors.Is(err, io.EOF) {
					break
				}
				fmt.Printf("error: %s\n", err.Error())
				return
			}

			data := buffer[:n]

			if i := bytes.IndexByte(data, '\n'); i != -1 {
				line := currentLine + string(data[:i])
				ch <- line

				currentLine = ""
				data = data[i+1:]

			}

			currentLine += string(data)
		}
	}()

	return ch
}
