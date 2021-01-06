package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	list := make([][][]int, 4)

	for i := 0; i < 4; i++ {
		list[i] = make([][]int, 3)

		for j := 0; j < 3; j++ {
			list[i][j] = make([]int, 10)
		}
	}

	for i := 0; i < a; i++ {
		var b, f, r, v int

		fmt.Scan(&b, &f, &r, &v)

		list[b-1][f-1][r-1] += v

	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			for n := 0; n < 10; n++ {
				fmt.Printf(" %d", list[i][j][n])
			}
			fmt.Printf("\n")
		}

		if i != 3 {
			fmt.Printf("####################\n")

		}

	}

}
