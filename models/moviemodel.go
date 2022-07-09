package models

type Movie struct {
	Name          string  `json:"name"`
	Rating        float32 `json:"rating"`
	YearRelease   int     `json:"yearRelease"`
	IsRecommended bool    `json:"isRecommended"`
	Id            int     `json:"id"`
}
