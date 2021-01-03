package main

import (
	"fmt"
	"sort"
)

func main() {
	var a int
	var S, H, C, D []int

	fmt.Scan(&a)

	for i := 0; i < a; i++ {
		var x int
		var y string

		fmt.Scan(&y, &x)

		if y == "S" {
			S = append(S, x)
			continue
		} else if y == "H" {
			H = append(H, x)
			continue
		} else if y == "C" {
			C = append(C, x)
			continue
		} else if y == "D" {
			D = append(D, x)
			continue
		}
	}
	sort.Ints(S)

	fmt.Println(S)
}

func check(a int) bool {
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
}
