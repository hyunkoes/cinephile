package api

import (
	"fmt"
	"time"

	"cinephile/modules/storage"

	"github.com/gin-gonic/gin"
)

type Movie struct {
	Movie_id       int       `json:"movie_id"`
	Is_adult       bool      `json:"is_adult"`
	Original_title string    `json:"original_title"`
	Kr_title       string    `json:"kr_title"`
	Poster_path    string    `json:"poster_path"`
	Release_date   time.Time `json:"release_date"`
	Overview       string    `json:"overview"`
}

func GetMovies(c *gin.Context) ([]Movie, error) {
	db := storage.DB()
	var length int
	_ = db.QueryRow(`select count(*) from movie`).Scan(&length)
	rows, _ := db.Query(`select * from movie`)
	defer rows.Close()
	fmt.Println(length)
	var mov Movie
	var movies []Movie
	for rows.Next() {
		err := rows.Scan(&mov.Movie_id, &mov.Is_adult, &mov.Original_title,
			&mov.Kr_title, &mov.Poster_path, &mov.Release_date, &mov.Overview)
		if err != nil {
			return []Movie{}, err
		}
		movies = append(movies, mov)
	}
	return movies, nil
}

func GetMovie(c *gin.Context) (Movie, error) {
	db := storage.DB()
	var length int
	_ = db.QueryRow(`select count(*) from movie`).Scan(&length)
	rows, _ := db.Query(`select * from movie`)
	defer rows.Close()
	fmt.Println(length)
	var mov Movie
	for rows.Next() {
		err := rows.Scan(&mov.Movie_id, &mov.Is_adult, &mov.Original_title,
			&mov.Kr_title, &mov.Poster_path, &mov.Release_date, &mov.Overview)
		if err != nil {
			return Movie{}, err
		}
	}
	return mov, nil
}
