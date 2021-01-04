package main

import (
	"fmt"
	"sort"
)

func main() {
	var a int
	var S [][]int

	fmt.Scan(&a)

	for i := 0; i < a; i++ {
		var x int
		var y string

		fmt.Scan(&y, &x)

	}
	sort.Ints(S)

	fmt.Println(S)
}

func check(a int, list []int) []int {
	sort.Ints(list)
	var out []int

	for i := 0; i < a; {
		l := list[i]

		if l != i+1 {
			out = append(out, l)
		} else {
			i++
		}
	}

	return out
}
