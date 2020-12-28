package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	x = x * x * x
	fmt.Println(x)
}
