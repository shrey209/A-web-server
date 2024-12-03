package main

import (
	"fmt"
	"net"
)

func main() {

	port := ":8080"
	listner, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("some error occured")
		return
	}
	defer listner.Close()

	for {
		conn, err := listner.Accept()
		if err != nil {
			fmt.Print("error occured during estalblishing the conection")
			continue
		}
		go handleReqHandler(conn)

	}

}
