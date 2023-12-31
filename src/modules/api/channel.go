package api

import (
	"cinephile/modules/storage"
	"errors"

	. "cinephile/modules/dto"
	. "cinephile/modules/tmdb"

	"github.com/gin-gonic/gin"
)

func GetChannels(c *gin.Context) ([]Channel, error) {
	return []Channel{}, nil
}

func GetChannel(c *gin.Context) (Channel, error) {
	channel_id, valid := c.GetQuery(`channel_id`)
	if !valid {
		return Channel{}, errors.New("Invalid channel id")
	}
	query := `
	SELECT 
		c.*,
		m.*
	FROM
		channel as c
	LEFT JOIN
		movie as m
	ON
		c.movie_id = m.movie_id
	WHERE
		channel_id = ` + channel_id + `
	`
	var channel Channel
	_ = channel
	db := storage.DB()
	row := db.QueryRow(query)
	err := row.Scan(&channel.Channel_id, &channel.Movie.Movie_id, &channel.Thread_count, &channel.Subscribe_count, &channel.Like_count, &channel.Movie.Movie_id, &channel.Movie.Is_adult, &channel.Movie.Original_title, &channel.Movie.Kr_title, &channel.Movie.Poster_path, &channel.Movie.Release_date, &channel.Movie.Overview)
	_ = err
	channel.Movie.Poster_path = TmdbPosterAPI(channel.Movie.Poster_path)
	return channel, nil
}

/*
Return subscribe request is success or not
Used in : Subscribe specific channel
*/
func SubscribeChannel(c *gin.Context) error {
	// db := storage.DB()
	// user := c.GetHeader(`user`)
	// channel_id, valid := c.GetQuery(`channel_id`)
	// if !valid {
	// 	return errors.New("Invalid channel_id")
	// }
	// var length int
	// result, err = db.Query(`select count(*) from user_subscribe`)
	// i := len(sql.Rows)
	return nil
}
