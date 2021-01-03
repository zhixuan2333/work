package main

import "fmt"

func main() {
	for i := 1; ; i++ {
		var x, y int
		fmt.Scan(&x)
		fmt.Scan(&y)
		if x == 0 && y == 0 {
			break
		}
		if x < y {
			fmt.Println(x, y)
		} else {
			fmt.Println(y, x)
		}
	}
}
