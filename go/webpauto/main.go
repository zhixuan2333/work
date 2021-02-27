package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

const src = "images\\"

// File info write to json
type File struct {
	Name string `json:"name"`
	Dir  string `json:"dir"`
	Md5  string `json:"md5"`
}

// Files file slince
type Files []File

func main() {

	check()
	now := open()
	WalkDir(now)

}

// WalkDir range file
func WalkDir(now Files) {
	var list File

	err := filepath.Walk(src,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				Md5 := HashFileMd5(path)
				if !Checkmd5(Md5, now) {

					if path[len(path)-5:] != ".webp" && path[len(path)-5:] != ".json" {
						list.Name = info.Name()
						list.Dir = path
						list.Md5 = Md5
						now = append(now, list)
						fmt.Println("-> " + path)
						Command(path, path+".webp")
					}

				}
			}

			return nil
		})
	if err != nil {
		log.Printf("Filepath Walk Failed: %e", err)
	}
	fmt.Println(now)
	Adddb(now)
}

//Command usr cwebp to Create webp file
func Command(src, dst string) []byte {
	out, err := exec.Command("cwebp", src, "-o", dst).Output()
	if err != nil {
		log.Printf("Failed: %s\n", err.Error())
	}

	return out
}

// Checkmd5 I don't know
func Checkmd5(md5 string, files Files) bool {
	for i := 0; i < len(files); i++ {
		if md5 == files[i].Md5 {
			return true
		}
	}
	return false
}

// Adddb add file to history.json
func Adddb(list Files) {

	f, _ := os.OpenFile(src+"history.json", os.O_CREATE|os.O_WRONLY, 0)
	enc := json.NewEncoder(f)

	err := enc.Encode(list)
	if err != nil {
		log.Printf("Error in encoding json: %e", err)
	}
}

func check() {
	if !Exists(src + "history.json") {
		f, err := os.Create(src + "history.json")
		if err != nil {
			log.Printf("Create history.json file err: %e", err)
		}
		defer f.Close()

	}

}

// Exists if the file Exist
func Exists(path string) bool {
	_, err := os.Stat(path)
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
		log.Printf("Check file md5 failed: %e", err)
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

func open() Files {
	jsonFile, err := os.Open(src + "history.json")

	if err != nil {
		log.Printf("Open history.json failed: %e", err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var files Files
	json.Unmarshal([]byte(byteValue), &files)

	fmt.Println(files)

	for i := 0; i < len(files); i++ {

		if !Exists(files[i].Dir) || !Exists(files[i].Dir+".webp") {
			files = append(files[:i], files[i+1:]...)
			i--
		}
	}
	fmt.Println(files)
	return files
}
