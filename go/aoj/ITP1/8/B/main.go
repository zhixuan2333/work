package main

import (
	"fmt"
)

func main() {
	for {
		var a string
		fmt.Scan(&a)
		if a == "0" {
			break
		}

		var ans int
		for _, k := range a {
			ans += int(k - '0')
		}
		fmt.Printf("%d\n", ans)
	}
}
