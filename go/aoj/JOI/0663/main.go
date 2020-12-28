package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	list := []int{a, b, c}

	st, nd := 0, 0

	for _, l := range list {
		if l == 1 {
			st++
		} else if l == 2 {
			nd++
		}
	}

	if st > nd {
		fmt.Println("1")
	} else {
		fmt.Println("2")
	}
}
