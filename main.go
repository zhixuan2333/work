package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

const src = "./"

// File info write to json
type File struct {
	Name string
	Dir  string
	Md5  string
}

func main() {

	check()
	WalkDir()

	// for _, l := range list {
	// 	out := Command(l, l+".webp")
	// 	fmt.Println(out)
	// }

	// check()
}

// WalkDir range file
func WalkDir() File {
	var list File

	err := filepath.Walk(src,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && path[len(path)-5:] != ".webp" {
				list.Name = info.Name()
				// list.Name = "test"
				list.Dir = src + path
				list.Md5 = HashFileMd5(path)

			}
			fmt.Println(list)

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

func check() f {
	if !Exists(src + "history.json") {
		f, err := os.Create(src + "history.json")
		if err != nil {
			log.Printf("Create history.json file err: %e", err)
		}
		defer f.Close()
	} else {
		f, err := os.Open(src + "history.json")
		if err != nil {
			log.Printf("Open history.json failed: %e", err)
		}
		defer f.Close()
	}
	return f

}

// Exists if the file Exist
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// HashFileMd5 Get file's md5
func HashFileMd5(filePath string) (md5Str string) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	hash := md5.New()
	if _, err = io.Copy(hash, file); err != nil {
		return
	}
	hashInBytes := hash.Sum(nil)[:16]
	md5Str = hex.EncodeToString(hashInBytes)
	return
}
