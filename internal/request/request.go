package request

import (
	"errors"
	"io"
	"strings"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	requestLine, err := parseRequestLine(data)
	if err != nil {
		return nil, err
	}

	request := &Request{
		*requestLine,
	}

	return request, nil
}

func parseRequestLine(data []byte) (*RequestLine, error) {
	dataString := string(data)
	parts := strings.Split(dataString, "\r\n")

	requestLineParts := strings.Split(parts[0], " ")
	if len(requestLineParts) != 3 {
		return nil, errors.New("invalid request line")
	}
	httpVersion := strings.Split(requestLineParts[2], "/")[1]

	requestLine := RequestLine{
		HttpVersion:   httpVersion,
		RequestTarget: requestLineParts[1],
		Method:        requestLineParts[0],
	}

	return &requestLine, nil
}
