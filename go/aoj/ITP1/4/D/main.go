package main

import (
	"fmt"
	"sort"
)

func main() {
	var a int
	sum := 0
	var list []int

	fmt.Scan(&a)

	for i := 1; i <= a; i++ {
		var x int

		fmt.Scan(&x)
		sum += x
		list = append(list, x)
	}

	sort.Ints(list)

	fmt.Printf("%d %d %d\n", list[0], list[len(list)-1], sum)
}
