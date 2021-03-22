package main

import (
	"fmt"
	"sort"
)

func main() {
	var A []int
	var B []int

	for i := 0; i < 4; i++ {
		var x int
		fmt.Scan(&x)

		A = append(A, x)
	}
	for i := 0; i < 2; i++ {
		var x int
		fmt.Scan(&x)

		B = append(B, x)
	}

	sort.Ints(A)
	sort.Ints(B)

	fmt.Println(A[1] + A[2] + A[3] + B[1])
}
