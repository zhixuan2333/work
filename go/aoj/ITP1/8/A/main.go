package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	instr := scanner.Text()

	for _, runVal := range instr {
		if runVal >= 'a' && runVal <= 'z' {
			fmt.Print(string(runVal - 'a' + 'A'))
		} else if runVal >= 'A' && runVal <= 'Z' {
			fmt.Print(string(runVal - 'A' + 'a'))
		} else {
			fmt.Print(string(runVal))
		}
	}
	fmt.Printf("\n")
}
