package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	hour := n / (60 * 60)
	n = n % (60 * 60)
	minute := n / 60
	second := n % 60
	fmt.Printf("%d:%d:%d\n", hour, minute, second)
}
