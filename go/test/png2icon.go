package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()
	fmt.Println(t.Format("2006_01_02"))
}
