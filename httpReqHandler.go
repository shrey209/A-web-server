package main

import (
	// Import your package (adjust "project" to your actual project folder name)
	"A-WEB-SERVER/http"
	"fmt"
	"net"
)

func handleReqHandler(conn net.Conn) {
	defer conn.Close()

	// Call the ParseHttpRequest function
	req, err := http.ParseHttpRequest(conn)
	if err != nil {
		fmt.Println("Error parsing HTTP request:", err)
		response := "HTTP/1.1 400 Bad Request\r\n\r\n"
		conn.Write([]byte(response))
		return
	}

	// Print the parsed HTTP request
	fmt.Println("HTTP Request Details:")
	fmt.Printf("Method: %s\n", req.Method)
	fmt.Printf("Path: %s\n", req.Path)
	fmt.Printf("Version: %s\n", req.Version)
	fmt.Printf("Headers: %v\n", req.Headers)
	fmt.Printf("Body: %s\n", req.Body)

	// Send a sample response
	response := "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\n\r\nRequest Parsed Successfully!"
	conn.Write([]byte(response))
}
