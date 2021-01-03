package main

import "fmt"

func main() {
	for true {
		var a, b int
		fmt.Scan(&a, &b)
		if a == 0 && b == 0 {
			break
		}
		h := 0
		for i := 0; i < a; i++ {
			w := 0
			for i := 0; i < b; i++ {
				if h == 0 || h == a-1 || w == 0 || w == b-1 {
					fmt.Printf("#")
				} else {
					fmt.Printf(".")
				}

				w++
			}
			fmt.Printf("\n")
			h++
		}
		fmt.Printf("\n")
	}
}
