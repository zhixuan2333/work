package main

import (
	"fmt"
)

func main() {
	for true {
		var a, b int
		fmt.Scan(&a, &b)
		if a == 0 && b == 0 {
			break
		}
		for i := 0; i < a; i++ {
			for i := 0; i < b; i++ {
				fmt.Printf("#")
			}
			fmt.Printf("\n")
		}
		fmt.Printf("\n")
	}
}
