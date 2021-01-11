package main

import "fmt"

func main() {
	var r, c int
	fmt.Scan(&r, &c)
	list := make([]int, c+1)

	for i := 0; i < r; i++ {
		var x int
		for j := 0; j < c; j++ {
			var a int
			fmt.Scan(&a)
			x += a

			fmt.Printf("%d ", a)
			list[j] += a
		}
		fmt.Printf("%d\n", x)
		list[c] += x

	}

	for i := 0; i < c; i++ {
		fmt.Printf("%d ", list[i])
	}
	fmt.Printf("%d\n", list[c])
}
