package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)
	var A []int
	for i := 0; i < N; i++ {
		var x int
		fmt.Scan(&x)
		A = append(A, x)
	}
	ans(A, N)

	for i := 1; i < N-1; i++ {
		key := A[i]
		j := i - 1
		for {
			if j >= 0 && A[j] > key {
				A[j+1] = A[j]
				j--
			} else {
				break
			}
		}
		A[j+1] = key
		ans(A, N)
	}
	if N != 1 {

		sort.Ints(A)
		ans(A, N)
	}

}

func ans(A []int, N int) {
	for i := 0; i < N; i++ {
		if i != N-1 {
			fmt.Printf("%d ", A[i])

		} else {
			fmt.Printf("%d\n", A[i])
		}
	}
}
