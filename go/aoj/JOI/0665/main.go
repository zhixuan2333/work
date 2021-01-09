package main

import (
	"fmt"
	"sort"
)

func main() {
	var a, b int
	var list []int
	fmt.Scan(&a, &b)

	for i := 0; i < a; i++ {
		var x int
		fmt.Scan(&x)

		list = append(list, x)
	}
	for i := 0; i < b; i++ {
		var x int
		fmt.Scan(&x)

		list = append(list, x)
	}

	sort.Ints(list)

	for _, k := range list {
		fmt.Printf("%d\n", k)
	}
}
