package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	var ans int
	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)

		if isPrime(x) {
			ans++
		}
	}
	fmt.Println(ans)
}

func isPrime(num int) bool {

	if 2 == num || 3 == num {
		return true
	}

	if num%6 != 1 && num%6 != 5 {
		return false
	}
	mid := int(math.Sqrt(float64(num)))

	for i := 5; i <= mid; i += 6 {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
	}

	return true
}
