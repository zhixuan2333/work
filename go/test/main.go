package main

import (
	"fmt"

	"github.com/gookit/color"
)

func main() {
	color.Cyan.Printf("Simple to use %s\n", "color")
	cyan := color.Red.Render
	fmt.Printf("hello %s\n", cyan("World"))
}
