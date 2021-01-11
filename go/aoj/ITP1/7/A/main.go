package main

import (
	"fmt"
)

func main() {
	for {
		var a, b, c int
		fmt.Scan(&a, &b, &c)

		if a == -1 && b == -1 && c == -1 {
			break
		}

		if a == -1 || b == -1 {
			fmt.Println("F")
			continue
		}

		num := a + b

		switch {
		case num >= 80:
			fmt.Println("A")

		case num >= 65:
			fmt.Println("B")

		case num >= 50:
			fmt.Println("C")

		case num >= 30:
			if c >= 50 {
				fmt.Println("C")
			} else {

				fmt.Println("D")
			}

		case num < 30:
			fmt.Println("F")

		}
	}
}
