package main

import "fmt"

func main() {
	for i := 1; ; i++ {
		var a int
		fmt.Scan(&a)
		if a == 0 {
			break
		}
		fmt.Printf("Case %d: %d\n", i, a)
	}
}
