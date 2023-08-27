package api

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

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
func SearchMovie(c *gin.Context) ([]MovieSearch, error) {
	db := storage.DB()
	search, valid := c.GetQuery(`keyword`)
	if !valid {
		return []MovieSearch{}, errors.New("No search query string")
	}
	query := `
	SELECT 
		m.*,
		c.channel_id
	FROM 
		movie as m
	LEFT JOIN
		channel as c
	ON
		m.movie_id = c.movie_id
	WHERE 
		original_title REGEXP '` + search + `'
		or
		kr_title REGEXP '` + search + `'
		;
	`
	rows, _ := db.Query(query)
	defer rows.Close()
	movies := make([]MovieSearch, 0)
	var mov MovieSearch
	var original_title sql.NullString
	var kr_title sql.NullString
	var poster_path sql.NullString
	var release_date sql.NullTime
	var overview sql.NullString

	for rows.Next() {
		err := rows.Scan(&mov.Movie_id, &mov.Is_adult, &original_title,
			&kr_title, &poster_path, &release_date, &overview, &mov.Channel_id)
		if err != nil {
			return []MovieSearch{}, err
		}
		if !original_title.Valid {
			mov.Original_title = ""
		} else {
			mov.Original_title = original_title.String
		}
		if !kr_title.Valid {
			mov.Kr_title = ""
		} else {
			mov.Kr_title = kr_title.String
		}
		if !poster_path.Valid {
			mov.Poster_path = ""
		} else {
			mov.Poster_path = poster_path.String
		}
		if !release_date.Valid {
			mov.Release_date = time.Time{}
		} else {
			mov.Release_date = release_date.Time
		}
		if !overview.Valid {
			mov.Overview = ""
		} else {
			mov.Overview = overview.String
		}

		movies = append(movies, mov)
	}
	return movies, nil
}
