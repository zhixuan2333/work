package main

import "fmt"

func main() {
	for {
		go test()
	}
}

func test() {
	for i := 1; ; i++ {
		fmt.Printf("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$\n")
	}
}
