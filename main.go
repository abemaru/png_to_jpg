package main

import (
	"fmt"
	"log"
	"os"
	"io/ioutil"
	"strings"
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

func main() {
	files, err := GetFilesInDir("sample")
	if err != nil {
		log.Fatal(err)
	}
	pngs := ExtractPNGImages(files)
	
}

