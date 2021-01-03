package main

import "fmt"

func main() {
	var w, h, x, y, r int
	fmt.Scan(&w, &h, &x, &y, &r)
	if x+r <= w && x-r >= 0 {
		if y-r >= 0 && y+r <= h {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}

	} else {
		fmt.Println("No")
	}

}
