package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
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
		inputDir string
	)

	flag.StringVar(&inputDir, "i", "../sounds/ps_saysounds_2019_0118/sound/misc/", "Input directory")
	flag.Parse()

	var sounds []string

	if err := filepath.Walk(inputDir, func(p string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		ext := filepath.Ext(p)
		if !isValidExt(ext) {
			return nil
		}
		sounds = append(sounds, filepath.Base(p))
		return nil
	}); err != nil {
		panic(err)
	}

	for i := range sounds {
		var cnt int
		for j := range sounds {
			if sounds[i] == sounds[j] {
				cnt++
			}
		}
		if cnt > 1 {
			fmt.Println(sounds[i], cnt)
		}
	}
}
