package main

import (
	"fmt"
)

func main() {
	for {
		var a, b int
		fmt.Scan(&a, &b)

		if a == 0 && b == 0 {
			break
		}

		var x int
		var list []int
		for i := 0; i < a; i++ {
			list = append(list, i+1)
		}

		for _, i := range list {
			for _, j := range list {
				if i == j {
					break
				}
				for _, n := range list {
					if n == i || n == j {
						break
					}
					if i+j+n == b {
						x++
					}
				}
			}
		}

		fmt.Println(x)
	}
}
