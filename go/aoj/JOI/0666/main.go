package main

import (
	"fmt"
	"sort"
)

func main() {
	var list []int

	for i := 0; i < 3; i++ {
		var x int
		fmt.Scan(&x)

		list = append(list, x)
	}

	sort.Ints(list)

	ans := list[1] + list[2]

	fmt.Printf("%d\n", ans)
}
