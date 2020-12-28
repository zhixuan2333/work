package main

import (
	"fmt"
	"sort"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	list := []int{a, b, c}

	sort.Ints(list)

	fmt.Printf("%d %d %d\n", list[0], list[1], list[2])

}
