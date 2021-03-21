package main

import "fmt"

func main() {
	var X, L, R int
	fmt.Scan(&X, &L, &R)

	if L <= X && X <= R {
		fmt.Println(X)
	} else if L > X {
		fmt.Println(L)
	} else {
		fmt.Println(R)
	}

}
