package main

import "fmt"

func main() {
	var a int
	var list []int
	fmt.Scan(&a)

	for i := 0; i < a; i++ {
		var x int
		fmt.Scan(&x)

		list = append(list, x)
	}

	for i := 0; i < a-1; i++ {

		fmt.Printf("%d ", list[a-i-1])

	}
	fmt.Printf("%d\n", list[0])
}
