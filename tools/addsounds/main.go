package main

import (
	"context"
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
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

type localSound struct {
	CmdName  string
	FileName string
	Path     string
	Ext      string
}

func isOnFirebase(ctx context.Context, bucket *storage.BucketHandle, cmdName, fileName string) (bool, error) {
	h := md5.New()
	io.WriteString(h, cmdName)
	p := filepath.Join("sounds", fmt.Sprintf("%x", h.Sum(nil))[:2], fileName)
	_, err := bucket.Object(p).Attrs(ctx)
	if err == storage.ErrObjectNotExist {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func main() {

	var (
		bucketURL string
		credPath  string
		inputDir  string
		outputDir string
		dryRun    bool
	)
	flag.StringVar(&bucketURL, "b", "", "URL of the storage bucket")
	flag.StringVar(&credPath, "c", "../../backend/config/firebase/firebase.json", "Path for file of firebase credential")
	flag.StringVar(&inputDir, "i", "../../sounds/saysound/ps_saysounds_2019_0118/sound/misc/", "Input directory")
	flag.StringVar(&outputDir, "o", "../../sounds/saysound/output", "Output directory")
	flag.BoolVar(&dryRun, "d", false, "Dry run")
	flag.Parse()

	if bucketURL == "" {
		log.Fatalln("bucketURL must be set")
	}

	ctx := context.Background()
	config := &firebase.Config{
		StorageBucket: bucketURL,
	}
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
	var dupIdx int
	for k, v := range sounds {
		if len(v) > 1 {
			for _, cmd := range v[1:] {
				var newName string
				lastWord := cmd.CmdName[len(cmd.CmdName)-1]
				index, err := strconv.Atoi(string(lastWord))
				if err != nil {
					index = 1
					for {
						newName = cmd.CmdName + strconv.Itoa(index)
						onFirebase, err := isOnFirebase(ctx, bucket, newName, newName+cmd.Ext)
						if err != nil {
							panic(err)
						}
						if _, ok := newSounds[newName]; ok || onFirebase {
							index++
							continue
						}
						break
					}
				} else {
					for {
						index++
						newName = cmd.CmdName[:len(cmd.CmdName)-1] + strconv.Itoa(index)
						onFirebase, err := isOnFirebase(ctx, bucket, newName, newName+cmd.Ext)
						if err != nil {
							panic(err)
						}
						if _, ok := newSounds[newName]; ok || onFirebase {
							continue
						}
						break
					}
				}
				prevPath := cmd.Path
				cmd.FileName = newName + cmd.Ext
				cmd.CmdName = newName
				cmd.Path = filepath.Join(filepath.Dir(cmd.Path), cmd.FileName)
				newSounds[cmd.CmdName] = []localSound{cmd}
				log.Printf("%d: %s->%s", dupIdx+1, prevPath, cmd.Path)
				dupIdx++
				if !dryRun {
					if err := os.Rename(prevPath, cmd.Path); err != nil {
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

	// Move to output directory
	for _, v := range newSounds {
		src := v[0].Path
		h := md5.New()
		io.WriteString(h, v[0].CmdName)
		hashed := fmt.Sprintf("%x", h.Sum(nil))
		dir := hashed[:2]
		dirPath := filepath.Join(outputDir, dir)
		dst := filepath.Join(dirPath, v[0].FileName)
		log.Printf("%s->%s", src, dst)
		if !dryRun {
			if err := os.MkdirAll(dirPath, 0755); err != nil {
				panic(err)
			}
			if err := os.Rename(src, dst); err != nil {
				panic(err)
			}
		}
	}
}
