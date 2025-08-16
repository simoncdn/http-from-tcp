package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
)

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
