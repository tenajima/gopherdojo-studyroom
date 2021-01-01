package main

import (
	"flag"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"path/filepath"

	"github.com/tenajima/kadai1/src"
)

func main() {
	fromFormat := flag.String("from", "jpeg", "変換前の形式")
	toFormat := flag.String("to", "png", "変換後の形式")

	flag.Parse()
	dirName := flag.Args()[0]
	outputDirName := filepath.Join(dirName, "convertedImages")
	if err := os.Mkdir(outputDirName, 0777); err != nil {
		panic(err)
	}

	err := filepath.Walk(dirName, func(path string, info os.FileInfo, err error) error {
		if ext := filepath.Ext(path); ext != ".jpg" && ext != ".png" {
			return nil
		}
		if err := src.ConvertImage(path, *fromFormat, *toFormat, outputDirName); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
