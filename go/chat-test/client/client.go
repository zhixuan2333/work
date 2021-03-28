package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:50000")
	if err != nil {
		log.Printf("Error dialing: %v\n", err)
		conn.Close()
		return
	}
	defer conn.Close()
	fmt.Println("First, what is your name?")
	var clientName string
	fmt.Scan(&clientName)

	fmt.Println("Type Q to quit.")
	for {
		fmt.Printf(">>> ")
		inputReader := bufio.NewReader(os.Stdin)
		input, _ := inputReader.ReadString('\n')
		input = strings.Trim(input, "\r\n")

		if input == "Q" {
			conn.Close()
			return
		}

		_, err = conn.Write([]byte(clientName + ": " + input))
	}
}
