package parsers

import (
	"encoding/json"
	"os"

	"github.com/takost/laughing-octo-robot/models"
)

func ParseContent(filePath string) ([]models.Movie, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var movies []models.Movie
	err = json.NewDecoder(file).Decode(&movies)
	if err != nil {
		return nil, err
	}

	return movies, nil
}
