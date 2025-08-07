package main

import (
	"fmt"
	"log"
	"net"

	"github.com/andycostintoma/httpfromtcp/internal/request"
)

const port = ":42069"

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error listening for TCP traffic: %s\n", err.Error())
	}
	defer listener.Close()

	fmt.Println("Listening for TCP traffic on", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("error: %s\n", err.Error())
		}
		fmt.Println("Accepted connection from", conn.RemoteAddr())

		request, err := request.RequestFromReader(conn)

		if err != nil {
			log.Fatalf("error: %s\n", err.Error())
		}

		reqLine := request.RequestLine
		fmt.Printf(`Request line:
- Method: %s
- Target: %s
- Version: %s`, reqLine.Method, reqLine.RequestTarget, reqLine.HttpVersion)

		// linesChan := getLinesChannel(conn)

		// for line := range linesChan {
		// 	fmt.Println(line)
		// }
		// fmt.Println("Connection to ", conn.RemoteAddr(), "closed")
	}
}
