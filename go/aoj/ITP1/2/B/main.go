package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if 0 <= a && a < b && b < c && c <= 100 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
