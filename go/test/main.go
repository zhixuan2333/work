package main

import "os"

func main() {

	os.Mkdir("world", os.ModePerm)
	os.Create("world/hello.go")
}
