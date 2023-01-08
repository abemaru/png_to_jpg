package main

import (
	"fmt"
	"log"
	"os"
	"io/ioutil"
	"strings"
	"path/filepath"
	_ "image/jpeg"
	"image/png"
	"image"
)


func GetFilesInDir(dirname string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}

	return files, err
}

func ExtractPNGImages(files []os.FileInfo) ([]os.FileInfo) {
	var pngs []os.FileInfo

	for _, file := range files {
		if strings.HasSuffix(file.Name(), "jpg") {
			pngs = append(pngs, file)
		}
	}
	return pngs
}

func CreateFileName(filepath string) (outputFilepath string) {
	return strings.Replace(filepath, "png", "jpg", -1)
}

func ConvertPNGtoJPG(dirname string,  pngs []os.FileInfo) () {
	for _, pngfile := range pngs {
		filepath := filepath.Join(dirname, pngfile.Name())
		file, err := os.Open(filepath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		img, _, err := image.Decode(file)
		if err != nil {
			log.Fatal(err)
		}

		outputPath := CreateFileName(filepath)
		out, err := os.Create(outputPath)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		fmt.Println(outputPath)
		png.Encode(out, img)
	}
}

func main() {
	files, err := GetFilesInDir("sample")
	if err != nil {
		log.Fatal(err)
	}
	pngs := ExtractPNGImages(files)
	ConvertPNGtoJPG("sample", pngs)
	
}

