package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

const src = "./images"

func main() {

	list := WalkDir()

	for _, l := range list {
		out := Command(l, l+".webp")
		fmt.Println(out)
	}

}

// WalkDir range file
func WalkDir() []string {
	var list []string
	err := filepath.Walk(src,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && path[len(path)-5:] != ".webp" {
				list = append(list, path)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	return list
}

//Command usr cwebp to Create webp file
func Command(src, dst string) []byte {
	// Print Go Version
	out, err := exec.Command("cwebp", src, "-o", dst).Output()
	if err != nil {
		log.Printf("Failed: %s\n", err.Error())
	}

	return out
}
