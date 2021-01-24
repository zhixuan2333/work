package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	ans := gcd(x, y)
	fmt.Println(ans)
}

func gcd(x, y int) int {
	tmp := x % y
	if tmp > 0 {
		return gcd(y, tmp)
	}
	return y

}
