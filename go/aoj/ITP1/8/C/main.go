package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)
	alpha := make([]int, 26)

	for {
		sentence, err := scanner.ReadString('\n')
		if err == io.EOF {
			break
		}
		for i := 0; i < len(sentence); i++ {
			if sentence[i] >= 'A' && sentence[i] <= 'Z' {
				alpha[sentence[i]-'A']++
			}
			if sentence[i] >= 'a' && sentence[i] <= 'z' {
				alpha[sentence[i]-'a']++
			}
		}

	}

	for i := 0; i < 26; i++ {
		fmt.Printf("%c : %d\n", i+'a', alpha[i])
	}
}
