package main

import (
	"context"
	"flag"
	"log"
	"path/filepath"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/lon9/SaySoundCloud/backend/models"
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
		bucketURL  string
		credPath   string
		dbURL      string
		dbProvider string
	)

	flag.StringVar(&bucketURL, "b", "", "URL of the storage bucket")
	flag.StringVar(&credPath, "c", "../../backend/config/firebase/firebase.json", "Path for file of firebase credential")
	flag.StringVar(&dbURL, "u", "../../backend/database/dev.db", "Database url")
	flag.StringVar(&dbProvider, "p", "sqlite3", "Database provider. sqlite3 or postgres or mysql")
	flag.Parse()

	if bucketURL == "" {
		log.Fatalln("bucketURL must be set")
	}
	config := &firebase.Config{
		StorageBucket: bucketURL,
	}
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, config, option.WithCredentialsFile(credPath))
	if err != nil {
		panic(err)
	}

	client, err := app.Storage(ctx)
	if err != nil {
		panic(err)
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(dbProvider, dbURL)
	if err != nil {
		panic(err)
	}
	defer db.Close()

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
		sound := &models.Sound{
			Name: strings.ReplaceAll(filepath.Base(objectAttr.Name), ext, ""),
			Path: filepath.Base(filepath.Dir(objectAttr.Name)) + "/" + filepath.Base(objectAttr.Name),
		}
		if err := db.Create(sound).Error; err != nil {
			log.Println(err)
		}
	}
}
