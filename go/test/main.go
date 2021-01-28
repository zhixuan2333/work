package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

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
	// 'bucket' is an object defined in the cloud.google.com/go/storage package.
	// See https://godoc.org/cloud.google.com/go/storage#BucketHandle
	// for more details.
	// [END cloud_storage_golang]

	log.Printf("Created bucket handle: %v\n", bucket)
}
