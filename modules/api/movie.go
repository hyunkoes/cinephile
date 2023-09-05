package api

import (
	"cinephile/modules/storage"
	. "cinephile/modules/tmdb"
	"database/sql"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

/*
Return 8 hottest movie ordered by count of thread and like during a week.
Used in : Movie search flow
*/
func GetHotMovies(c *gin.Context) ([]MovieSearch, error) {
	db := storage.DB()
	_ = db
	query := `
SELECT
    m.movie_id,
    m.is_adult,
    m.original_title,
    m.kr_title,
    m.poster_path,
    m.release_date,
    m.overview,
    c.channel_id,
    GROUP_CONCAT(DISTINCT g.genre_id),
    GROUP_CONCAT(DISTINCT g.genre_name)
FROM
    movie AS m
LEFT JOIN
    genre_relation AS gr ON m.movie_id = gr.movie_id
LEFT JOIN
    genre AS g ON gr.genre_id = g.genre_id
LEFT JOIN
    channel AS c ON m.movie_id = c.movie_id
LEFT JOIN
    thread AS t ON c.channel_id = t.channel_id
LEFT JOIN
    thread_recommend AS tr ON t.thread_id = tr.thread_id
WHERE
    DATE(t.created_at) >= DATE(NOW() - INTERVAL 1 WEEK)
GROUP BY
    m.movie_id, c.channel_id
ORDER BY
    COUNT(DISTINCT t.thread_id) DESC,
    SUM(tr.is_recommended) DESC
LIMIT
    8;
`
	rows, err := db.Query(query)
	if err != nil {
		return []MovieSearch{}, err
	}

	defer rows.Close()
	movies := make([]MovieSearch, 0)
	var mov MovieSearch
	var original_title sql.NullString
	var kr_title sql.NullString
	var poster_path sql.NullString
	var release_date sql.NullTime
	var overview sql.NullString
	var genres_ids string
	var genres_names string
	for rows.Next() {
		err := rows.Scan(&mov.Movie_id, &mov.Is_adult, &original_title,
			&kr_title, &poster_path, &release_date, &overview, &mov.Channel_id, &genres_ids, &genres_names)
		genres := make([]Genre, 0)

		g_ids := strings.Split(genres_ids, ",")
		g_names := strings.Split(genres_names, ",")
		for i := 0; i < len(g_ids); i++ {
			id, _ := strconv.Atoi(g_ids[i])
			genres = append(genres, Genre{Genre_id: id, Genre_name: g_names[i]})
		}
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
		mov.Genres = genres
		mov.Poster_path = TmdbPosterAPI(mov.Poster_path)
		movies = append(movies, mov)
	}
	return movies, nil
}

/*
Return all movies
Used in : Test
*/
func GetMovies(c *gin.Context) ([]Movie, error) {
	db := storage.DB()
	var length int
	_ = db.QueryRow(`select count(*) from movie`).Scan(&length)
	rows, _ := db.Query(`select * from movie`)
	defer rows.Close()
	var mov Movie
	var movies []Movie
	for rows.Next() {
		err := rows.Scan(&mov.Movie_id, &mov.Is_adult, &mov.Original_title,
			&mov.Kr_title, &mov.Poster_path, &mov.Release_date, &mov.Overview)
		if err != nil {
			return []Movie{}, err
		}
		mov.Poster_path = TmdbPosterAPI(mov.Poster_path)
		movies = append(movies, mov)
	}
	return movies, nil
}

/*
Return detail movie info
Used in : Single movie page
*/
func GetMovie(c *gin.Context) (Movie, error) {
	db := storage.DB()
	var length int
	_ = db.QueryRow(`select count(*) from movie`).Scan(&length)
	rows, _ := db.Query(`select * from movie`)
	defer rows.Close()
	var mov Movie
	for rows.Next() {
		err := rows.Scan(&mov.Movie_id, &mov.Is_adult, &mov.Original_title,
			&mov.Kr_title, &mov.Poster_path, &mov.Release_date, &mov.Overview)
		if err != nil {
			return Movie{}, err
		}
		mov.Poster_path = TmdbPosterAPI(mov.Poster_path)
	}
	return mov, nil
}

/*
Return 8 movies filtered by keyword and cursor. ( cursor pagination )
Used in : Movie search flow
*/
func SearchMovie(c *gin.Context) ([]MovieSearch, int, error) {
	db := storage.DB()
	search, valid := c.GetQuery(`keyword`)
	if !valid {
		return []MovieSearch{}, 0, errors.New("No search query string")
	}
	cursor, valid := c.GetQuery(`cursor`)
	if !valid {
		cursor = "0"
	}

	query := `
	SELECT 
		m.movie_id,
		m.is_adult,
		m.original_title,
		m.kr_title,
		m.poster_path,
		m.release_date,
		m.overview,
		c.channel_id,
		GROUP_CONCAT(g.genre_id) AS genre_ids,
    	GROUP_CONCAT(g.genre_name) AS genre_names
	FROM
		movie m
	LEFT JOIN
		channel c ON m.movie_id = c.movie_id
	LEFT JOIN
		genre_relation gr ON m.movie_id = gr.movie_id
	LEFT JOIN
		genre g ON gr.genre_id = g.genre_id
	WHERE 
		(original_title REGEXP '` + search + `'
		or
		kr_title REGEXP '` + search + `')
		and m.movie_id > ` + cursor + `
	GROUP BY
		m.movie_id, m.original_title, m.kr_title, c.channel_id
	LIMIT
		8
		;
	`
	rows, _ := db.Query(query)
	defer rows.Close()
	movies := make([]MovieSearch, 0)

	// @ Todo : make simple below code.
	var mov MovieSearch
	var original_title sql.NullString
	var kr_title sql.NullString
	var poster_path sql.NullString
	var release_date sql.NullTime
	var overview sql.NullString
	var lastCursor int
	var genres_ids string
	var genres_names string
	for rows.Next() {
		err := rows.Scan(&mov.Movie_id, &mov.Is_adult, &original_title,
			&kr_title, &poster_path, &release_date, &overview, &mov.Channel_id, &genres_ids, &genres_names)
		genres := make([]Genre, 0)

		g_ids := strings.Split(genres_ids, ",")
		g_names := strings.Split(genres_names, ",")
		for i := 0; i < len(g_ids); i++ {
			id, _ := strconv.Atoi(g_ids[i])
			genres = append(genres, Genre{Genre_id: id, Genre_name: g_names[i]})
		}
		if err != nil {
			return []MovieSearch{}, 0, err
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
		mov.Genres = genres
		mov.Poster_path = TmdbPosterAPI(mov.Poster_path)
		lastCursor = mov.Movie_id
		movies = append(movies, mov)
	}
	return movies, lastCursor, nil
}
