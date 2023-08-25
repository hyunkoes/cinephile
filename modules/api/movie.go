package api

import (
	"errors"
	"fmt"

	"cinephile/modules/storage"

	"github.com/gin-gonic/gin"
)

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
func SearchMovie(c *gin.Context) ([]Movie, error) {
	db := storage.DB()
	search, valid := c.GetQuery(`search`)
	if !valid {
		return []Movie{}, errors.New("No search query string")
	}
	query := `
	SELECT 
		*
	FROM 
		movie
	WHERE 
		original_title REGEXP '` + search + `'
		or
		kr_title REGEXP '` + search + `'
		;
	`
	rows, _ := db.Query(query)
	defer rows.Close()
	movies := make([]Movie, 0)
	var mov Movie
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
