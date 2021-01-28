package main

import (
	"context"
	"io"
	"io/ioutil"
	"log"
<<<<<<< Updated upstream
=======
	"os"
>>>>>>> Stashed changes

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func main() {
<<<<<<< Updated upstream

	config := &firebase.Config{
		StorageBucket: "test-fb724.appspot.com",
	}
	opt := option.WithCredentialsFile("test-fb724-firebase-adminsdk-dcryi-81e7333440.json")
	app, err := firebase.NewApp(context.Background(), config, opt)
=======

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

	downloadFile("history.json", bucket)

}

// uploadFile upload an object.
func uploadFile(firename string, bucket *storage.BucketHandle) {

	contentType := "text/plain"
	ctx := context.Background()

	writer := bucket.Object(firename).NewWriter(ctx)
	writer.ObjectAttrs.ContentType = contentType
	writer.ObjectAttrs.CacheControl = "no-cache"
	writer.ObjectAttrs.ACL = []storage.ACLRule{
		{
			Entity: storage.AllUsers,
			Role:   storage.RoleReader,
		},
	}

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
>>>>>>> Stashed changes
	if err != nil {
		log.Fatalln(err)
	}

<<<<<<< Updated upstream
	client, err := app.Storage(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		log.Fatalln(err)
	}
	// 'bucket' is an object defined in the cloud.google.com/go/storage package.
	// See https://godoc.org/cloud.google.com/go/storage#BucketHandle
	// for more details.
	// [END cloud_storage_golang]

	log.Printf("Created bucket handle: %v\n", bucket)
=======
	f, err := os.Create(filename)
	if err != nil {
		log.Printf("Create download file failed: %e", err)
	}
	defer f.Close()

	f.Write(data)

	log.Printf("Downloaded contents: %v\n", string(data))
>>>>>>> Stashed changes
}
