package main

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
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

var src = "./"

func main() {

	var option string
	flag.StringVar(&option, "o", "init", "option init or finish")

	if option == "finish" {
		normalSync(1)

	} else {

		initSync()
		go Sync()
		// mcstart()

		// normalSync(1)

		mux := http.NewServeMux()
		mux.HandleFunc("/", Root)

		port := os.Getenv("PORT")
		if port == "" {
			port = "12345"
		}

		http.ListenAndServe(":"+port, mux)
	}

}

// Root send message to root
func Root(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("main.html")
		if err != nil {
			log.Printf("root failed: %e", err)
		}
		t.Execute(w, nil)
	} else {
		w.Write([]byte("Error"))
	}
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

	cmd := exec.Command("java", "-Xms200M", "-Xmx400", "-jar", "server.jar", "nogui")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	fmt.Println("Result: " + out.String())

	// cmd := exec.Command("java", "-jar", "server.jar")

	// stdin, _ := cmd.StdinPipe()
	// stdout, _ := cmd.StdoutPipe()

	// if err := cmd.Start(); err != nil {
	// 	fmt.Println("Execute failed when Start:" + err.Error())
	// 	return
	// }

	// // stdin.Write([]byte("go text for grep\n"))
	// // stdin.Write([]byte("go test text for grep\n"))
	// stdin.Close()

	// outbytes, _ := ioutil.ReadAll(stdout)
	// stdout.Close()

	// if err := cmd.Wait(); err != nil {
	// 	fmt.Println("Execute failed when Wait:" + err.Error())
	// 	return
	// }

	// fmt.Println("Execute finished:" + string(outbytes))

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

	wg := sync.WaitGroup{}
	var files Files
	var tmp string

	err = filepath.Walk(src,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.Name() != "main.go" && info.Name() != "__debug_bin" && info.Name() != "init.json" && info.Name() != "token.json" && info.Name() != "Dockerfile" && info.Name() != "entrypoint.sh" && info.Name() != "heroku.yml" && info.Name() != "server.jar" {
				Md5 := HashFileMd5(path)

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
				tmp = strings.Replace(path, string(filepath.Separator), "/", -1)

				if !info.IsDir() {
					wg.Add(1)
					go func(firename string, bucket *storage.BucketHandle, wg *sync.WaitGroup) {
						uploadFile(firename, bucket)
						wg.Done()
					}(tmp, bucket, &wg)

				}

				list := File{info.Name(), tmp, Md5}
				files = append(files, list)
				// }
				// !info.IsDir() &&

			}

			return nil
		})
	if err != nil {
		log.Printf("Filepath Walk Failed: %v\n", err)
	}

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
		log.Printf("Open init.json failed: %v\n", err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var files Files
	json.Unmarshal([]byte(byteValue), &files)
	jsonFile.Close()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		resp, err := http.Get("https://launcher.mojang.com/v1/objects/f02f4473dbf152c23d7d484952121db0b36698cb/server.jar")
		if err != nil {
			log.Printf("download server.jar failed: %v\n", err)
		}
		defer resp.Body.Close()

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Create server.jar failed: %v\n", err)
		}
		ioutil.WriteFile("server.jar", data, 0644)
		wg.Done()
	}(&wg)

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
