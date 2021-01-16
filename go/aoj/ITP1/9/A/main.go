package main

import (
	"fmt"
	"strings"
)

func main() {
	var w string
	fmt.Scan(&w)

	var ans int
	for {
		var t string
		fmt.Scan(&t)

		if t == "END_OF_TEXT" {
			break
		}
		t = strings.ToLower(t)
		if t == w {
			ans++
		}
	}
	fmt.Println(ans)
}
