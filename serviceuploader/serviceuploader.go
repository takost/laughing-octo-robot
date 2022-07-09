package serviceuploader

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/takost/laughing-octo-robot/models"
)

func UploadMovies(movies []models.Movie) error {
	url := "http://localhost:4000/films"

	for _, movie := range movies {
		fmt.Printf("Uploading movie: %s\n", movie.Name)

		data, err := json.Marshal(movie)
		if err != nil {
			panic(err)
		}

		resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		switch resp.StatusCode {
		case http.StatusCreated:
			fmt.Printf("Movie %s uploaded successfully\n", movie.Name)
		case http.StatusBadRequest:
			fmt.Printf("Incorrect movie data: %s\n", string(data))
		default:
			return errors.New("Movie upload failed with status code: " + resp.Status)
		}
	}
	return nil
}
