package main

import "fmt"

func main() {
	var ring, shaped string
	fmt.Scan(&ring, &shaped)
	// Write a program which finds a text in a ring shaped pattern.
	// and return Yes or No.
	// The pattern and text are given as below.
	// Pattern: vanceknowledgetoad
	// Text: advance
	// Output: Yes

	for i := 0; i < len(ring); i++ {
		if ring[i] == shaped[0] {
			for j := 0; j < len(shaped); j++ {
				if ring[(i+j)%len(ring)] != shaped[j] {
					break
				}
				if j == len(shaped)-1 {
					fmt.Println("Yes")
					return
				}
			}
		}
	}
	fmt.Println("No")

}
