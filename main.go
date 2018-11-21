package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	tcpPort := os.Getenv("TCP_PORT")
	if tcpPort == "" {
		tcpPort = "6557"
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%s", tcpPort))
	if err != nil {
		log.Panicln(err)
	}

	message := fmt.Sprintf("Echo server listens on %s\n", tcpPort)

	log.Println("Listening on TCP port", tcpPort)
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Panicln(err)
		}

		log.Println("Established tcp connection")

		go handleTCPRequest(conn, message)
	}
}

func handleTCPRequest(conn net.Conn, message string) {
	defer conn.Close()
	defer log.Println("TCP connection closed.")

	_, err := conn.Write([]byte(message))
	if err != nil {
		return
	}

	for {
		buf := make([]byte, 1024)
		size, err := conn.Read(buf)
		if err != nil {
			return
		}
		data := string(buf[:size])

		log.Printf("Received Data (converted to string): %s", data)
		_, err = conn.Write(buf[:size])
		if err != nil {
			return
		}
	}
}
