package serviceuploader

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/takost/laughing-octo-robot/models"
)

const MaxRetries int = 5
const RetryInterval float32 = 0.5

func UploadMovies(movies []models.Movie) error {
	url := "http://localhost:4000/films"

	for _, movie := range movies {
		fmt.Printf("Uploading movie: %s\n", movie.Name)

		data, err := json.Marshal(movie)
		if err != nil {
			panic(err)
		}

		var retries int = 0
		var currentInterval float32 = RetryInterval
		for retries < MaxRetries {
			retries++
			resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			if resp.StatusCode == http.StatusCreated {
				fmt.Printf("Movie %s uploaded successfully\n", movie.Name)
				break
			} else if resp.StatusCode == http.StatusBadRequest {
				fmt.Printf("Incorrect movie data: %s\n", string(data))
				break
			} else {
				fmt.Printf("WARNING: Movie upload failed with status code: " + resp.Status)
				if retries < MaxRetries {
					fmt.Printf("Retrying in %.1f seconds\n", currentInterval)
					time.Sleep(time.Duration(currentInterval) * time.Second)
					currentInterval *= 2
				} else {
					return errors.New("ERROR: Movie upload failed with status code: " + resp.Status)
				}
			}
		}
	}
	return nil
}
