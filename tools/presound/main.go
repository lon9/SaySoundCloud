package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/lon9/soundboard/backend/models"
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

type localSound struct {
	CmdName  string
	FileName string
	Path     string
	Ext      string
}

func main() {

	var (
		inputDir   string
		outputDir  string
		dbURL      string
		dbProvider string
		dryRun     bool
	)

	flag.StringVar(&inputDir, "i", "../../sounds/saysound/ps_saysounds_2019_0118/sound/misc/", "Input directory")
	flag.StringVar(&outputDir, "o", "../../sounds/saysound/output", "Output directory")
	flag.StringVar(&dbURL, "u", "../../backend/database/dev.db", "Database url")
	flag.StringVar(&dbProvider, "p", "sqlite3", "Database provider. sqlite3 or postgres or mysql")
	flag.BoolVar(&dryRun, "d", false, "Dry run")
	flag.Parse()

	sounds := make(map[string][]localSound)
	newSounds := make(map[string][]localSound)

	// get files
	var numSounds int
	if err := filepath.Walk(inputDir, func(p string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		ext := filepath.Ext(p)
		if !isValidExt(ext) {
			return nil
		}
		sound := localSound{
			FileName: filepath.Base(p),
			CmdName:  strings.ReplaceAll(filepath.Base(p), ext, ""),
			Path:     p,
			Ext:      ext,
		}
		if _, ok := sounds[sound.CmdName]; ok {
			sounds[sound.CmdName] = append(sounds[sound.CmdName], sound)
			newSounds[sound.CmdName] = append(newSounds[sound.CmdName], sound)
		} else {
			sounds[sound.CmdName] = []localSound{sound}
			newSounds[sound.CmdName] = []localSound{sound}
		}
		numSounds++
		return nil
	}); err != nil {
		panic(err)
	}

	// resolve command name
	for k, v := range sounds {
		if len(v) > 1 {
			for i, cmd := range v[1:] {
				var newName string
				lastWord := cmd.CmdName[len(cmd.CmdName)-1]
				index, err := strconv.Atoi(string(lastWord))
				if err != nil {
					index = 1
					for {
						newName = cmd.CmdName + strconv.Itoa(index)
						if _, ok := newSounds[newName]; ok {
							index++
							continue
						}
						break
					}
				} else {
					for {
						index++
						newName = cmd.CmdName[:len(cmd.CmdName)-1] + strconv.Itoa(index)
						if _, ok := newSounds[newName]; ok {
							continue
						}
						break
					}
				}
				cmd.FileName = newName + cmd.Ext
				cmd.CmdName = newName
				cmd.Path = filepath.Join(filepath.Dir(cmd.Path), cmd.FileName)
				newSounds[cmd.CmdName] = []localSound{cmd}
				log.Printf("%s->%s", v[i].CmdName, cmd.CmdName)
				if !dryRun {
					if err := os.Rename(v[i].Path, cmd.Path); err != nil {
						panic(err)
					}
				}
			}
		}
		newSounds[k] = []localSound{v[0]}
	}

	// assertion
	for _, v := range newSounds {
		if len(v) != 1 {
			log.Fatal("invlid length")
		}
	}
	if len(newSounds) != numSounds {
		log.Fatalf("The number of sounds is invalid %d:%d", len(newSounds), numSounds)
	}

	db, err := gorm.Open(dbProvider, dbURL)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(new(models.Sound))

	// Move to output directory
	var idx int
	records := make([]models.Sound, len(newSounds))
	for _, v := range newSounds {
		src := v[0].Path
		h := md5.New()
		io.WriteString(h, v[0].CmdName)
		hashed := fmt.Sprintf("%x", h.Sum(nil))
		dir := hashed[:2]
		dirPath := filepath.Join(outputDir, dir)
		dst := filepath.Join(dirPath, v[0].FileName)
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			panic(err)
		}
		log.Printf("%s->%s", src, dst)
		if !dryRun {
			if err := os.Rename(src, dst); err != nil {
				panic(err)
			}
		}
		record := models.Sound{
			Name: v[0].CmdName,
			Path: filepath.Join(dir, v[0].FileName),
		}
		records[idx] = record
		idx++
	}

	if !dryRun {
		for i := range records {
			if err := db.Create(&records[i]).Error; err != nil {
				panic(err)
			}
		}
	}
}
