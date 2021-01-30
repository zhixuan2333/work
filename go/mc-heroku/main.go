package main

import (
	"context"
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
	"strings"
	"sync"
	"time"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// File info write to json
type File struct {
	Name string `json:"name"`
	Dir  string `json:"dir"`
	Md5  string `json:"md5"`
}

// Files file slince
type Files []File

var src = ".\\"

// Version minecraft version
const Version = "1.16.3"

func main() {

	initSync()
	go Sync()
	mcstart()

	// normalSync(1)
}

// Sync is mcserver
func Sync() {
	fmt.Println("Start minecraft server")

	for {
		time.Sleep(time.Second * 600)

		normalSync(0)

	}

}

func mcstart() {

	cmd := exec.Command("java", "-jar", "server.jar")

	stdin, _ := cmd.StdinPipe()
	stdout, _ := cmd.StdoutPipe()

	if err := cmd.Start(); err != nil {
		fmt.Println("Execute failed when Start:" + err.Error())
		return
	}

	// stdin.Write([]byte("go text for grep\n"))
	// stdin.Write([]byte("go test text for grep\n"))
	stdin.Close()

	outbytes, _ := ioutil.ReadAll(stdout)
	stdout.Close()

	if err := cmd.Wait(); err != nil {
		fmt.Println("Execute failed when Wait:" + err.Error())
		return
	}

	fmt.Println("Execute finished:" + string(outbytes))

	normalSync(1)

}

// normalSync sync mincraft server
func normalSync(status int) {

	fmt.Println("Start sync mc world")

	config := &firebase.Config{
		StorageBucket: "test-fb724.appspot.com",
	}
	opt := option.WithCredentialsFile("token.json")
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Storage(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		log.Fatalln(err)
	}

	jsonFile, err := os.Open("init.json")

	if err != nil {
		log.Printf("Open init.json failed: %e", err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var now Files
	json.Unmarshal([]byte(byteValue), &now)

	jsonFile.Close()
	wg := sync.WaitGroup{}
	var files Files

	err = filepath.Walk(src,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.Name() != "main.go" && info.Name() != "__debug_bin" && info.Name() != "init.json" && info.Name() != "token.json" && info.Name() != "Dockerfile" && info.Name() != "entrypoint.sh" {
				Md5 := HashFileMd5(path)
				if !Checkmd5(Md5, now) || Md5 == "" {

					if status == 0 && info.Name() == "session.lock" {
						return nil

					}

					// if path[len(path)-5:] != ".webp" || path[len(path)-5:] != ".json" {
					// 	list.Name = info.Name()
					// 	list.Dir = path
					// 	list.Md5 = Md5
					// 	now = append(now, list)
					// 	fmt.Println("NOW")

					// }
					// wg.Add(1)
					// go func(firename string, bucket *storage.BucketHandle, wg sync.WaitGroup) {
					// 	uploadFile(firename, bucket)
					// 	wg.Done()
					// }(path, bucket, wg)
					tmp := strings.Replace(path, string(filepath.Separator), "/", -1)

					list := File{info.Name(), tmp, Md5}
					files = append(files, list)
					if !info.IsDir() {
						wg.Add(1)
						go func(firename string, bucket *storage.BucketHandle, wg *sync.WaitGroup) {
							uploadFile(firename, bucket)
							wg.Done()
						}(tmp, bucket, &wg)

					}
				}
				// }
				// !info.IsDir() &&

			}

			return nil
		})
	if err != nil {
		log.Printf("Filepath Walk Failed: %v\n", err)
	}
	fmt.Println(now)
	Adddb(files)
	wg.Wait()

	// wg.Wait()
	uploadFile("init.json", bucket)

}

// initSync init minecraft server
func initSync() Files {
	config := &firebase.Config{
		StorageBucket: "test-fb724.appspot.com",
	}
	opt := option.WithCredentialsFile("token.json")
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Storage(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		log.Fatalln(err)
	}

	downloadFile("init.json", bucket)

	jsonFile, err := os.Open("init.json")

	if err != nil {
		log.Printf("Open init.json failed: %e", err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var files Files
	json.Unmarshal([]byte(byteValue), &files)
	jsonFile.Close()

	wg := sync.WaitGroup{}

	for _, v := range files {
		if v.Dir == ".\\" {
			continue
		}
		if v.Md5 == "" {
			os.Mkdir(v.Dir, os.ModePerm)

			continue
		}
		fmt.Println(v.Dir)
		wg.Add(1)

		go func(filename string, bucket *storage.BucketHandle, wg *sync.WaitGroup) {
			downloadFile(filename, bucket)
			wg.Done()
		}(v.Dir, bucket, &wg)
		// go downloadFile(v.Dir, bucket)
	}
	wg.Wait()

	return files
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

// uploadFile upload an object.
func uploadFile(filename string, bucket *storage.BucketHandle) {

	ctx := context.Background()

	writer := bucket.Object(filename).NewWriter(ctx)

	fmt.Println(filename)
	f, err := os.Open(filename)
	if _, err = io.Copy(writer, f); err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	if err := writer.Close(); err != nil {
		log.Fatalln(err)
	}
}

func downloadFile(filename string, bucket *storage.BucketHandle) {
	ctx := context.Background()
	rc, err := bucket.Object(filename).NewReader(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer rc.Close()

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		log.Fatalln(err)
	}

	f, err := os.Create(filename)
	if err != nil {
		log.Printf("Create download file failed: %v\n", err)
	}
	defer f.Close()

	f.Write(data)
}

// Adddb add file to history.json
func Adddb(list Files) {

	if !Exists("init.json") {
		_, err := os.Create("init.json")
		if err != nil {
			log.Printf("Create init.json failed: %v\n", err)
		}

	} else {
		err := os.Remove("init.json")
		if err != nil {
			log.Printf("Remove init.json failed: %v\n", err)
		}
		_, err = os.Create("init.json")
		if err != nil {
			log.Printf("Create init.json failed: %v\n", err)
		}
	}

	f, err := os.OpenFile("init.json", os.O_RDWR, 0644)
	if err != nil {
		log.Printf("Open init.json failed: %v\n", err)
	}
	defer f.Close()
	enc := json.NewEncoder(f)

	err = enc.Encode(list)
	if err != nil {
		log.Printf("Error in encoding json: %v\n", err)
	}
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

// Checkmd5 I don't know
func Checkmd5(md5 string, files Files) bool {
	for i := 0; i < len(files); i++ {
		if md5 == files[i].Md5 {
			return true
		}
	}
	return false
}
