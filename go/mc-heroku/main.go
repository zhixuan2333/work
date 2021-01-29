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
	config := &firebase.Config{
		StorageBucket: "test-fb724.appspot.com",
	}
	opt := option.WithCredentialsFile("test-fb724-firebase-adminsdk-dcryi-81e7333440.json")
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

	fmt.Println("hello")
	normalSync(Files{}, bucket)

	// go Sync()
	// mcstart()

}

// Sync is mcserver
func Sync() {
	fmt.Println("Sync now")

	for {
		time.Sleep(time.Second * 30)

		fmt.Println("-----------")

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

}

// normalSync sync mincraft server
func normalSync(now Files, bucket *storage.BucketHandle) {

	var files Files

	err := filepath.Walk(src,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && info.Name() != "main.go" {
				Md5 := HashFileMd5(path)
				// if !Checkmd5(Md5, now) {

				// 	// if path[len(path)-5:] != ".webp" || path[len(path)-5:] != ".json" {
				// 	// 	list.Name = info.Name()
				// 	// 	list.Dir = path
				// 	// 	list.Md5 = Md5
				// 	// 	now = append(now, list)
				// 	// 	fmt.Println("NOW")

				// 	// }

				// }
				list := File{info.Name(), path, Md5}
				files = append(files, list)
				uploadFile(path, bucket)
			}

			return nil
		})
	if err != nil {
		log.Printf("Filepath Walk Failed: %e", err)
	}
	fmt.Println(now)
	Adddb(files)

	uploadFile("init.json", bucket)

}

// initSync init minecraft server
func initSync() Files {
	config := &firebase.Config{
		StorageBucket: "test-fb724.appspot.com",
	}
	opt := option.WithCredentialsFile("test-fb724-firebase-adminsdk-dcryi-81e7333440.json")
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

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var files Files
	json.Unmarshal([]byte(byteValue), &files)

	for _, v := range files {
		downloadFile(v.Dir, bucket)
	}

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
func uploadFile(firename string, bucket *storage.BucketHandle) {

	ctx := context.Background()

	writer := bucket.Object(firename).NewWriter(ctx)

	f, err := os.Open(firename)
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
		log.Printf("Create download file failed: %e", err)
	}
	defer f.Close()

	f.Write(data)
}

// Adddb add file to history.json
func Adddb(list Files) {

	f, _ := os.OpenFile("init.json", os.O_CREATE|os.O_WRONLY, 0)
	defer f.Close()
	enc := json.NewEncoder(f)

	err := enc.Encode(list)
	if err != nil {
		log.Printf("Error in encoding json: %e", err)
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
