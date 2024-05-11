package main

import (
	"encoding/json"
	"fmt"
)

type Movies struct {
	Title       MoviesParts `json:"Title"`
	ReleaseDate int         `json:"Date"`
	ImdbRating  float32     `json:"Rating"`
}

type MoviesParts struct {
	Title string `json:"Title"`
	Date  int    `json:"Date"`
}

func main() {
	moviesParts := MoviesParts{Title: "SpiderMan 2", Date: 2006}
	moviesParts = MoviesParts{Title: "SpiderMan 2", Date: 2006}
	movies := Movies{Title: moviesParts, ReleaseDate: 2002, ImdbRating: 8.5}
	moviesJson, err := json.MarshalIndent(movies, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(moviesJson))
}
