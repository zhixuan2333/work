package main

import (
	"fmt"
)

func main() {
	var v int
	list := make([][]int, 4)
	mark := []string{"S", "H", "C", "D"}

	for i := 0; i < 4; i++ {
		list[i] = make([]int, 13)
	}

	fmt.Scan(&v)

	for i := 0; i < v; i++ {
		var mk string
		var num int

		fmt.Scan(&mk, &num)

		for j := 0; j < 4; j++ {
			if mark[j] == mk {
				list[j][num-1] = 1
			}
		}

	}

	for j := 0; j < 4; j++ {
		for i := 0; i < 13; i++ {
			if list[j][i] == 0 {
				fmt.Printf("%s %d\n", mark[j], i+1)
			}
		}
	}
}
