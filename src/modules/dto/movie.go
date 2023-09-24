package dto

import "time"

type Movie struct {
	Movie_id       int       `json:"movieId"`
	Is_adult       bool      `json:"isAdult"`
	Original_title string    `json:"originalTitle"`
	Kr_title       string    `json:"krTitle"`
	Poster_path    string    `json:"posterPath"`
	Release_date   time.Time `json:"releaseDate"`
	Overview       string    `json:"overview"`
	Genres         []Genre   `json:"genres"`
}
type MovieSearch struct {
	Movie_id       int       `json:"movieId"`
	Channel_id     int       `json:"channelId"`
	Is_adult       bool      `json:"isAdult"`
	Original_title string    `json:"originalTitle"`
	Kr_title       string    `json:"krTitle"`
	Poster_path    string    `json:"posterPath"`
	Release_date   time.Time `json:"releaseDate"`
	Overview       string    `json:"overview"`
	Genres         []Genre   `json:"genres"`
}
