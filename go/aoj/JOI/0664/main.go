package main

import (
	"fmt"
	"strings"
)

func main() {
	var a int
	var b string
	fmt.Scan(&a)
	fmt.Scan(&b)

	list := strings.Split(b, "")

	x := 0

	for _, y := range list {
		if inVowel(y) {
			x++
		}
	}

	fmt.Println(x)

}

func inVowel(c string) bool {
	vowel := []string{"a", "e", "i", "o", "u"}

	for _, v := range vowel {
		if v == c {
			return true
		}
	}
	return false
}
