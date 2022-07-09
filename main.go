package main

import (
	"fmt"

	"github.com/takost/laughing-octo-robot/parsers"
	"github.com/takost/laughing-octo-robot/serviceuploader"
)

func main() {
	filePaths := []string{
		"./test.json",
	}

	for _, filePath := range filePaths {
		movies, err := parsers.ParseContent(filePath)
		if err != nil {
			panic(err)
		}
		err = serviceuploader.UploadMovies(movies)
		if err != nil {
			fmt.Println(err)
		}
	}
}
