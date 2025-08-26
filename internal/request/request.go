package request

import (
	"fmt"
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
		return nil, fmt.Errorf("poorly formatted request-line: %s", parts)
	}

	method := requestLineParts[0]
	for _, c := range method {
		if c < 'A' || c > 'Z' {
			return nil, fmt.Errorf("invalid method: %s", method)
		}
	}

	requestTarget := requestLineParts[1]

	versionParts := strings.Split(requestLineParts[2], "/")

	httpPart := versionParts[0]
	if httpPart != "HTTP" {
		return nil, fmt.Errorf("unrecognized HTTP-version: %s", httpPart)
	}

	version := versionParts[1]
	if version != "1.1" {
		return nil, fmt.Errorf("unrecognized HTTP-version: %s", version)
	}

	requestLine := RequestLine{
		HttpVersion:   version,
		RequestTarget: requestTarget,
		Method:        method,
	}

	return &requestLine, nil
}
