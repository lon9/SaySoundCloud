package main

import (
	"context"
	"flag"
	"log"
	"path/filepath"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

var exts = []string{".mp3", ".wav", ".ogg", ".aac"}

func isValidExt(ext string) bool {
	for _, e := range exts {
		if ext == e {
			return true
		}
	}
	return false
}

func main() {

	var (
		bucketURL string
		credPath  string
	)

	flag.StringVar(&bucketURL, "b", "", "URL of the storage bucket")
	flag.StringVar(&credPath, "c", "../../backend/config/firebase/firebase.json", "Path for file of firebase credential")
	flag.Parse()

	if bucketURL == "" {
		log.Fatalln("bucketURL must be set")
	}
	config := &firebase.Config{
		StorageBucket: bucketURL,
	}
	opt := option.WithCredentialsFile(credPath)
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, config, opt)
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

	it := bucket.Objects(ctx, nil)
	for {
		objectAttr, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Println(err)
			continue
		}
		ext := filepath.Ext(objectAttr.Name)
		if !isValidExt(ext) {
			continue
		}
		log.Printf("Updating %s\n", objectAttr.Name)
		obj := bucket.Object(objectAttr.Name)
		if _, err := obj.Update(
			ctx,
			storage.ObjectAttrsToUpdate{
				CacheControl: "public, max-age=86400",
			},
		); err != nil {
			panic(err)
		}
	}
}
