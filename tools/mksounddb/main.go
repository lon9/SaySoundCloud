package main

import (
	"flag"
	"os"
	"path/filepath"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/lon9/SaySoundCloud/backend/models"
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
		soundDir   string
		dbURL      string
		dbProvider string
		dryRun     bool
	)

	flag.StringVar(&soundDir, "i", "../../sounds/saysound/output", "Sound directory")
	flag.StringVar(&dbURL, "u", "../../backend/database/dev.db", "Database url")
	flag.StringVar(&dbProvider, "p", "sqlite3", "Database provider. sqlite3 or postgres or mysql")
	flag.BoolVar(&dryRun, "d", false, "Dry run")
	flag.Parse()

	var sounds []models.Sound
	if err := filepath.Walk(soundDir, func(p string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		ext := filepath.Ext(p)
		if !isValidExt(ext) {
			return nil
		}
		base := filepath.Base(p)
		sound := models.Sound{
			Name: strings.ReplaceAll(base, ext, ""),
			Path: filepath.Join(filepath.Base(filepath.Dir(p)), base),
		}
		sounds = append(sounds, sound)
		return nil
	}); err != nil {
		panic(err)
	}

	db, err := gorm.Open(dbProvider, dbURL)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(new(models.Sound))
	if !dryRun {
		for i := range sounds {
			if err := db.Create(&sounds[i]).Error; err != nil {
				panic(err)
			}
		}
	}
}
