package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./messages.txt")
	if err != nil {
		fmt.Println("couln't open this file: messages.txt", err)
	}
	defer file.Close()

	currentLine := ""

	for {
		buffer := make([]byte, 8)

		n, err := file.Read(buffer)
		if err != nil {
			break
		}

		data := buffer[:n]

		if i := bytes.IndexByte(data, '\n'); i != -1 {
			line := currentLine + string(data[:i])
			fmt.Printf("read: %s\n", line)

			currentLine = ""
			data = data[i+1:]
		}

		currentLine += string(data)
	}
}
