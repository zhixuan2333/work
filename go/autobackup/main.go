package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Zip 打包成zip文件
func Zip(srcdir string, output string) {

	// 创建：zip文件
	zipfile, err := os.Create(output)
	if err != nil {
		log.Printf("Create zip file failed: %s\n", err.Error())
	}
	defer zipfile.Close()

	// 打开：zip文件
	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	// 遍历路径信息
	filepath.Walk(srcdir, func(path string, info os.FileInfo, _ error) error {

		// 如果是源路径，提前进行下一个遍历
		if path == srcdir {
			return nil
		}

		// 获取：文件头信息
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			log.Printf("Get file header failed: %s\n", err.Error())
		}
		header.Name = strings.TrimPrefix(path, srcdir+`\`)

		// 判断：文件是不是文件夹
		if info.IsDir() {
			header.Name += `/`
		} else {
			// 设置：zip的文件压缩算法
			header.Method = zip.Deflate
		}

		// 创建：压缩包头部信息
		writer, err := archive.CreateHeader(header)
		if err != nil {
			log.Printf("Create zip header failed: %s\n", err.Error())
		}
		if !info.IsDir() {
			file, _ := os.Open(path)
			defer file.Close()
			io.Copy(writer, file)
		}
		return nil
	})
}

func gettime() string {
	t := time.Now()
	t1 := t.Format("2006_01_02")
	return t1
}

func main() {
	t := gettime()
	dst := "D:\\workspace" + t + ".zip"
	src := "D:\\workspace"
	Zip(src, dst)
	fmt.Println("finish")
}
