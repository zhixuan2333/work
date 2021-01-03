package main

import "fmt"

func main() {
	for true {
		var a, b int
		var c string
		fmt.Scan(&a, &c, &b)

		if c == "?" {
			break
		} else if c == "+" {
			fmt.Println(a + b)

		} else if c == "-" {
			fmt.Println(a - b)

		} else if c == "*" {
			fmt.Println(a * b)

		} else if c == "/" {
			fmt.Println(a / b)

		}
	}
}
