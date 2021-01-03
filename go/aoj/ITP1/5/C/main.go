package main

import "fmt"

func main() {
	for true {
		var a, b int
		fmt.Scan(&a, &b)
		if a == 0 && b == 0 {
			break
		}
		for h := 0; h < a; h++ {
			for w := 0; w < b; w++ {
				if (h+w)%2 == 0 {
					fmt.Printf("#")
					continue
				} else {
					fmt.Printf(".")
					continue
				}
			}
			fmt.Printf("\n")
		}
		fmt.Printf("\n")
	}
}
