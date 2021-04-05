package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	name = "workspace" + time.Now().Format("2006_01_02") + ".zip"
	src  = "C:\\Users\\jinzh\\OneDrive\\Document\\workspace\\"
	dst  = "C:\\local\\backup\\" + name
)

// compression: compression file
func compression(src, zipsrc string) {

	// creat: zip file
	zipfile, err := os.Create(zipsrc)
	if err != nil {
		log.Fatalf("creat zip file err :%s/n", err.Error())
		return
	}
	defer zipfile.Close()

	// open: zip file
	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	// file遍历路径信息
	filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			fDest, err := archive.Create(path[len(src)+0:])
			if err != nil {
				log.Printf("Create failed: %s\n", err.Error())
				return nil
			}

			fSrc, err := os.Open(path)
			if err != nil {
				log.Printf("Open failed: %s\n", err.Error())
				return nil
			}
			defer fSrc.Close()
			_, err = io.Copy(fDest, fSrc)
			if err != nil {
				log.Printf("Copy failed: %s\n", err.Error())
				return nil
			}
		}
		return nil
	})
}

func main() {
	fmt.Println("Auto backup program")
	fmt.Println("filename:" + name)
	compression(src, dst)
	fmt.Println("Finish!")
}
