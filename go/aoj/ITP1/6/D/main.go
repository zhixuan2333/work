package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	var b [100]int
	var A [100][100]int

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			var x int
			fmt.Scan(&x)
			A[i][j] = x
		}
	}

	for i := 0; i < m; i++ {
		var x int
		fmt.Scan(&x)
		b[i] = x
	}

	for i := 0; i < n; i++ {
		var x int
		for j := 0; j < m; j++ {
			x += A[i][j] * b[j]
		}
		fmt.Println(x)
	}

}
