package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64
	fmt.Scan(&r)

	p := 2 * math.Pi * r
	s := math.Pi * r * r

	fmt.Printf("%f %f", s, p)

}
