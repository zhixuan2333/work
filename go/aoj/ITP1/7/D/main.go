package main

import "fmt"

func main() {
	var n, m, l int
	fmt.Scan(&n, &m, &l)

	var A, B [100][100]int

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			var x int
			fmt.Scan(&x)

			A[i][j] = x
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < l; j++ {
			var x int
			fmt.Scan(&x)

			B[i][j] = x
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < l; j++ {
			var ans int
			for k := 0; k < m; k++ {
				ans += A[i][k] * B[k][j]
			}
			if j == l-1 {
				fmt.Printf("%d", ans)
			} else {
				fmt.Printf("%d ", ans)
			}
		}
		fmt.Printf("\n")
	}
}
