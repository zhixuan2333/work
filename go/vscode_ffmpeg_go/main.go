package main

import (
	"archive/zip"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"sync"

	"github.com/vbauerster/mpb"
	"github.com/vbauerster/mpb/decor"
)

//vscodeV get vscode version
func vscodeV() []byte {
	// Print Go Version
	out, err := exec.Command("code", "-v").Output()
	if err != nil {
		log.Printf("Get vscode version failed: %s\n", err.Error())
	}
	re := []byte(out)
	return re
}

// atLine get string by int
func atLine(f []byte, n int) (s string) {
	r := bytes.NewReader(f)
	bufReader := bufio.NewReader(r)
	for i := 1; ; i++ {
		line, _, err := bufReader.ReadLine()
		if err != nil {
			break
		}
		if i == n {
			s = string(line)
			break
		}
	}
	return s
}

// electron get version
func electron(version string) string {
	// get .yarnrc file
	r, err := http.Get("https://raw.githubusercontent.com/Microsoft/vscode/" + version + "/.yarnrc")
	if err != nil {
		log.Printf("get electron version failed: %s", err.Error())
	}
	defer r.Body.Close()

	// get .yarnrc version
	b, err := ioutil.ReadAll(r.Body)
	s := string(b)

	// re match version
	rule, err := regexp.Compile(`".*?"`)
	if err != nil {
		log.Printf("re rule is failed: %s\n", err.Error())
	}
	results := rule.FindAllString(s, -1)
	i := results[1]
	end := len(i) - 1
	re := i[1:end]

	return re
}

// get system OS info
func systemversion() string {
	OS := runtime.GOOS
	if OS == "windows" {
		return "win32"
	}
	if OS == "darwin" {
		return "darwin"
	}
	return "linux"

}

// Resource Resource
type Resource struct {
	Filename string
	Url      string
}

// Downloader Downloader
type Downloader struct {
	wg         *sync.WaitGroup
	pool       chan *Resource
	Concurrent int
	HttpClient http.Client
	TargetDir  string
	Resources  []Resource
}

// NewDownloader NewDownloader
func NewDownloader(targetDir string) *Downloader {
	concurrent := runtime.NumCPU()
	return &Downloader{
		wg:         &sync.WaitGroup{},
		TargetDir:  targetDir,
		Concurrent: concurrent,
	}
}

// AppendResource AppendResource
func (d *Downloader) AppendResource(filename, url string) {
	d.Resources = append(d.Resources, Resource{
		Filename: filename,
		Url:      url,
	})
}

// Download main progress
func (d *Downloader) Download(resource Resource, progress *mpb.Progress) error {
	defer d.wg.Done()
	d.pool <- &resource
	finalPath := d.TargetDir + "/" + resource.Filename
	// 创建临时文件
	target, err := os.Create(finalPath + ".tmp")
	if err != nil {
		return err
	}

	// 开始下载
	req, err := http.NewRequest(http.MethodGet, resource.Url, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		target.Close()
		return err
	}
	defer resp.Body.Close()
	fileSize, _ := strconv.Atoi(resp.Header.Get("Content-Length"))
	// 创建一个进度条
	bar := progress.AddBar(
		int64(fileSize),
		// 进度条前的修饰
		mpb.PrependDecorators(
			decor.CountersKibiByte("% .2f / % .2f"), // 已下载数量
			decor.Percentage(decor.WCSyncSpace),     // 进度百分比
		),
		// 进度条后的修饰
		mpb.AppendDecorators(
			decor.EwmaETA(decor.ET_STYLE_GO, 90),
			decor.Name(" ] "),
			decor.EwmaSpeed(decor.UnitKiB, "% .2f", 60),
		),
	)
	reader := bar.ProxyReader(resp.Body)
	defer reader.Close()
	// 将下载的文件流拷贝到临时文件
	if _, err := io.Copy(target, reader); err != nil {
		target.Close()
		return err
	}

	// 关闭临时并修改临时文件为最终文件
	target.Close()
	if err := os.Rename(finalPath+".tmp", finalPath); err != nil {
		return err
	}
	<-d.pool
	return nil
}

// Start Download
func (d *Downloader) Start() error {
	d.pool = make(chan *Resource, d.Concurrent)
	fmt.Println("Start Download")
	p := mpb.New(mpb.WithWaitGroup(d.wg))
	for _, resource := range d.Resources {
		d.wg.Add(1)
		go d.Download(resource, p)
	}
	p.Wait()
	d.wg.Wait()
	return nil
}

// ffmpeg unzip file
func ffmpeg(dst, src, OS string) {
	zr, err := zip.OpenReader(src)
	defer zr.Close()
	if err != nil {
		log.Printf("Open zip file failed: %s\n", err.Error())
	}

	// Creat folder
	if dst != "" {
		if err := os.MkdirAll(dst, 0755); err != nil {
			log.Printf("Creat dst folder failed: %s\n", err.Error())
		}
	}

	// Traverse zr and write files to disk
	for _, file := range zr.File {
		if OS == "win32" {
			if file.Name == "ffmpeg.dll" {
				path := filepath.Join(dst, file.Name)
				// Get the Reader
				fr, err := file.Open()
				if err != nil {
					log.Printf("Open file failed: %s\n", err.Error())
				}

				// Create the Write corresponding to the file to be written
				fw, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, file.Mode())
				if err != nil {
					log.Printf("Write failed: %s\n", err.Error())
				}

				n, err := io.Copy(fw, fr)
				if err != nil {
					log.Printf("Write file failed: %s\n", err.Error())
				}

				// Output the decompressed result
				fmt.Printf("%s was successfully decompressed, and %d characters of data were written\n", path, n)

				fw.Close()
				fr.Close()
			}
		} else if OS == "darwin" {
			if file.Name == "Electron.app" || file.Name == "Contents" || file.Name == "Frameworks" || file.Name == "Electron Framework.framework" || file.Name == "Versions" || file.Name == "A" || file.Name == "Libraries" || file.Name == "libffmpeg.dylib" {
				path := filepath.Join(dst, file.Name)
				// Get the Reader
				fr, err := file.Open()
				if err != nil {
					log.Printf("Open file failed: %s\n", err.Error())
				}

				// Create the Write corresponding to the file to be written
				fw, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, file.Mode())
				if err != nil {
					log.Printf("Write failed: %s\n", err.Error())
				}

				n, err := io.Copy(fw, fr)
				if err != nil {
					log.Printf("Write file failed: %s\n", err.Error())
				}

				// Output the decompressed result
				fmt.Printf("%s was successfully decompressed, and %d characters of data were written\n", path, n)

				fw.Close()
				fr.Close()
			}
		} else if OS == "linux" {
			if file.Name == "libffmpeg.so" {
				path := filepath.Join(dst, file.Name)
				// Get the Reader
				fr, err := file.Open()
				if err != nil {
					log.Printf("Open file failed: %s\n", err.Error())
				}

				// Create the Write corresponding to the file to be written
				fw, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, file.Mode())
				if err != nil {
					log.Printf("Write failed: %s\n", err.Error())
				}

				n, err := io.Copy(fw, fr)
				if err != nil {
					log.Printf("Write file failed: %s\n", err.Error())
				}

				// Output the decompressed result
				fmt.Printf("%s was successfully decompressed, and %d characters of data were written\n", path, n)

				fw.Close()
				fr.Close()
			}
		}
	}
}

func main() {
	vscode := vscodeV()
	Vversion := atLine(vscode, 1)
	arch := atLine(vscode, 3)
	yarnrc := electron(Vversion)
	OS := systemversion()
	fmt.Printf("vscode: %s\n", Vversion)
	fmt.Printf("arch: %s\n", arch)
	fmt.Printf("version: %s\n", yarnrc)
	fmt.Printf("OS: %s\n", OS)
	url := "https://github.com/electron/electron/releases/download/v" + yarnrc + "/electron-v" + yarnrc + "-" + OS + "-" + arch + ".zip"
	filename := "electron-v" + yarnrc + "-" + OS + "-" + arch + ".zip"
	fmt.Println(url)
	downloader := NewDownloader("./")
	downloader.AppendResource(filename, url)
	err := downloader.Start()
	if err != nil {
		log.Printf("download failed: %s\n", err.Error())
	}
	ffmpeg("", filename, OS)
	os.Remove(filename)

}
