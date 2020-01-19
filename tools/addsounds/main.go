package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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
		inputDir string
	)

	flag.StringVar(&inputDir, "i", "../sounds/saysound/ps_saysounds_2019_0118/sound/misc/", "Input directory")
	flag.Parse()

	sounds := make(map[string][]localSound)

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
		} else {
			sounds[sound.CmdName] = []localSound{sound}
		}
		return nil
	}); err != nil {
		panic(err)
	}

	for _, v := range sounds {
		if len(v) > 1 {
			for _, cmd := range v[1:] {
				var newName string
				lastWord := cmd.CmdName[len(cmd.CmdName)-1]
				index, err := strconv.Atoi(string(lastWord))
				if err != nil {
					newName = cmd.CmdName + "1"
				} else {
					for {
						index++
						newName = cmd.CmdName[:len(cmd.CmdName)-1] + strconv.Itoa(index)
						if _, ok := sounds[newName]; ok {
							continue
						}
						break
					}
				}
				cmd.FileName = newName + cmd.Ext
				cmd.CmdName = newName
				cmd.Path = filepath.Join(filepath.Dir(cmd.Path), cmd.FileName)
				fmt.Println(cmd)
				sounds[cmd.CmdName] = []localSound{cmd}
			}
		}
	}
}
