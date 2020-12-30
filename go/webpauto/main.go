package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {

	list := WalkDir()

	fmt.Println(list)

}

// WalkDir range file
func WalkDir() []string {
	var list []string
	err := filepath.Walk(".\\images",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				list = append(list, path)
				fmt.Println(path)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	return list
}
