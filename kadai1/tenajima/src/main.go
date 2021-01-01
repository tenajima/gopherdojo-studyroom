package src

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

func ConvertImage(path string, inputFormat string, outputFormat string, outputDirName string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	img, format, err := image.Decode(f)
	iwt := imageWithExt{img, format}
	if inputFormat != iwt.format {
		fmt.Println("inputFormat: ", inputFormat)
		fmt.Println("iwt.format: ", iwt.format)
		return errors.New("指定したフォーマットと画像のフォーマットが一致してないよ")
	}
	if outputFormat == iwt.format {
		return errors.New("変換元と変換先が一緒だよ")
	}
	if err != nil {
		return err
	}
	baseName := strings.Split(filepath.Base(path), ".")[0]
	switch outputFormat {
	case "png":
		of, err := os.Create(filepath.Join(outputDirName, baseName+".png"))
		defer of.Close()
		if err != nil {
			return err
		}
		png.Encode(of, iwt.img)
	case "jpeg":
		of, err := os.Create(filepath.Join(outputDirName, baseName+".jpg"))
		defer of.Close()
		if err != nil {
			return err
		}
		jpeg.Encode(of, iwt.img, &jpeg.Options{Quality: 100})
	}
	return nil
}

type imageWithExt struct {
	img    image.Image
	format string
}
