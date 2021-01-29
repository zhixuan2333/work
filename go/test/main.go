package main

import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"os"

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

	log.Printf("Created bucket handle: %v\n", bucket)

	downloadFile("init.json", bucket)

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
