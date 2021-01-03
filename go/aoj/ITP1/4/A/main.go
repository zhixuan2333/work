package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	c := float64(a)
	d := float64(b)

	st := a / b
	sd := a % b
	fd := c / d

	fmt.Printf("%d %d %f", st, sd, fd)
}
