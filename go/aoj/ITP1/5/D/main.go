package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			fmt.Printf(" %d", i)
			continue
		}
		for j := i; j > 0; j = j / 10 {
			if j%10 == 3 {
				fmt.Printf(" %d", i)
				break
			}
		}
	}
	fmt.Println("")
	return
}
