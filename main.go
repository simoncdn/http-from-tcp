package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./messages.txt")
	if err != nil {
		fmt.Println("couln't open this file: messages.txt", err)
	}

	linesChan := getLinesChannel(file)

	for line := range linesChan {
		fmt.Printf("read: %s\n", line)
	}
}
