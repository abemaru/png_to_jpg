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

	for _, file := range files {
		if strings.HasSuffix(file.Name(), "jpg") {
			fmt.Println(file.Name())
		}
	}
	return files
}

func main() {
	files, err := GetFilesInDir("sample")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ExtractPNGImages(files))
}

