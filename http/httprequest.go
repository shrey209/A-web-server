package http

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
)

// HttpRequest represents the parsed HTTP request structure
type HttpRequest struct {
	Method  string
	Path    string
	Version string
	Headers map[string]string
	Body    string
}

// ParseHttpRequest parses the incoming connection to extract an HTTP request
func ParseHttpRequest(conn net.Conn) (*HttpRequest, error) {
	reader := bufio.NewReader(conn)

	// Parse the request line
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		return nil, errors.New("failed to read request line")
	}
	requestLine = strings.TrimSpace(requestLine)
	requestLineParts := strings.Split(requestLine, " ")
	if len(requestLineParts) != 3 {
		return nil, errors.New("malformed request line")
	}
	method, path, version := requestLineParts[0], requestLineParts[1], requestLineParts[2]

	// Parse headers
	headers := make(map[string]string)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return nil, errors.New("failed to read headers")
		}
		line = strings.TrimSpace(line)
		if line == "" { // End of headers
			break
		}
		headerParts := strings.SplitN(line, ": ", 2)
		if len(headerParts) == 2 {
			key := strings.ToLower(headerParts[0]) // Normalize header keys
			headers[key] = headerParts[1]
		} else {
			// Optionally log a warning for malformed headers
			fmt.Printf("Warning: Malformed header ignored: %s\n", line)
		}
	}

	// Parse body
	var body string
	if contentLength, ok := headers["content-length"]; ok {
		length, err := strconv.Atoi(contentLength)
		if err != nil {
			return nil, fmt.Errorf("invalid content-length: %s", contentLength)
		}
		bodyBuf := make([]byte, length)
		_, err = reader.Read(bodyBuf)
		if err != nil {
			return nil, fmt.Errorf("failed to read body: %v", err)
		}
		body = string(bodyBuf)
	}

	// Return the parsed request
	return &HttpRequest{
		Method:  method,
		Path:    path,
		Version: version,
		Headers: headers,
		Body:    body,
	}, nil
}
