package api

import (
	"time"
)

type Thread struct {
	Thread_id      int       `json:"threadId"`
	Movie_id       int       `json:"movieId"`
	Parent         int       `json:"parent,default=-1"`
	Content        string    `json:"content"`
	Created_at     time.Time `json:"createdAt"`
	Updated_at     time.Time `json:"updatedAt"`
	Is_recommended bool      `json:"isLiked"`
	Channel        Channel   `json:"channel"`
	Author         User      `json:"author"`
}

type Thread_detail struct {
	Self   Thread   `json:"self"`
	Parent Thread   `json:"parent"`
	Child  []Thread `json:"child"`
}

type Channel struct {
	Id             int    `json:"id"`
	Original_title string `json:"originalTitle"`
	Kr_title       string `json:"krTitle"`
	Poster         string `json:"poster"`
}
type Movie struct {
	Movie_id       int       `json:"movie_id"`
	Is_adult       bool      `json:"is_adult"`
	Original_title string    `json:"original_title"`
	Kr_title       string    `json:"kr_title"`
	Poster_path    string    `json:"poster_path"`
	Release_date   time.Time `json:"release_date"`
	Overview       string    `json:"overview"`
}

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Image int    `json:"image"`
}
type Thread_recommend struct {
	Thread_id      int       `json:"threadId"`
	Channel_id     int       `json:"channelId"`
	Original_title string    `json:"originalTitle"`
	Kr_title       string    `json:"krTitle"`
	Movie_id       int       `json:"movieId"`
	Email          string    `json:"email"`
	Parent         int       `json:"parent,default=-1"`
	Content        string    `json:"content"`
	Created_at     time.Time `json:"createdAt"`
	Updated_at     time.Time `json:"updatedAt"`
	Is_recommended bool      `json:"isRecommended"`
}
type RecommendForm struct {
	Thread_id int    `json:"threadId"`
	Email     string `json:"email"`
}
