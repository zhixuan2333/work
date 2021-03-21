package main

import "fmt"

func main() {
	var N, A, B, C, D int
	fmt.Scan(&N, &A, &B, &C, &D)

	var X, Y int

	for i := 1; ; i++ {
		if N <= A*i {
			X = B * i
			break
		}
	}
	for i := 1; ; i++ {
		if N <= C*i {
			Y = D * i
			break
		}
	}

	if X > Y {
		fmt.Println(Y)
	} else {
		fmt.Println(X)
	}
}
