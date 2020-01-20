package main

import (
	"context"
	"flag"
	"log"
	"path/filepath"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"github.com/jinzhu/gorm"
	"github.com/lon9/soundboard/backend/models"
	"google.golang.org/api/option"
)

func main() {

	var (
		bucketURL  string
		credPath   string
		dbURL      string
		dbProvider string
	)

	flag.StringVar(&bucketURL, "b", "", "URL of the storage bucket")
	flag.StringVar(&credPath, "c", "firebase.json", "Path for file of firebase credential")
	flag.StringVar(&dbURL, "u", "../../backend/database/dev.db", "Database url")
	flag.StringVar(&dbProvider, "p", "sqlite3", "Database provider. sqlite3 or postgres or mysql")
	flag.Parse()

	if bucketURL == "" {
		log.Fatalln("bucketURL must be set")
	}
	config := &firebase.Config{
		StorageBucket: bucketURL,
	}
	opt := option.WithCredentialsFile(credPath)
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

	db, err := gorm.Open(dbProvider, dbURL)
	if err != nil {
		panic(err)
	}

	var sounds []models.Sound
	if err := db.Find(&sounds).Error; err != nil {
		panic(err)
	}

	for _, sound := range sounds {
		dst := filepath.Join("sounds", sound.Path)

		ctx := context.Background()
		obj := bucket.Object(dst)

		// Set cache-control as 1 day
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
