package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	B := make([]int, M+1)

	for i := 0; i < N; i++ {
		var x int
		fmt.Scan(&x)

		B[x]++

	}

	sort.Ints(B)

	fmt.Println(B[M])

}
