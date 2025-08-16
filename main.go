package main

import (
	"fmt"
	"io"
	"os"
	"strings"
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
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("read error:", err)
			return
		}

		stringedBuffer := string(buffer[:n])
		parts := strings.Split(stringedBuffer, "\n")

		for i := 0; i < len(parts)-1; i++ {
			line := currentLine + parts[i]
			fmt.Printf("read: %s\n", line)

			currentLine = ""
		}

		currentLine += parts[len(parts)-1]
	}
}
