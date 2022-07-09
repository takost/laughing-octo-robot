package main

import (
	"github.com/takost/laughing-octo-robot/parsers"
	"github.com/takost/laughing-octo-robot/serviceuploader"
)

func main() {
	filePaths := []string{
		"/workspaces/laughing-octo-robot/test.json",
	}

	movies, err := parsers.ParseContent(filePaths[0])
	if err != nil {
		panic(err)
	}

	serviceuploader.UploadMovies(movies)
}
