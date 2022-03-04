package main

import "fmt"

func main() {
	a := make([]int, 26)

	for i := 0; i < 4; i++ {
		var n int
		fmt.Scan(&n)
		a[n-1]++
	}

	var count int
	for i := 0; i < 26; i++ {
		if a[i] == 2 {
			count++
		}
	}
	if count == 2 {
		fmt.Println("1")
	} else {
		fmt.Println("0")
	}
}
