package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("Start the Server ...")

	listener, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		log.Printf("Error Listening: %v\n", err)
		return

	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Printf("Error accepting: %v\n", err)
			return
		}
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			log.Printf("Error reading: %v\n", err)
			return
		}
		fmt.Printf("Received data: %v\n", string(buf[:len]))
	}
}
