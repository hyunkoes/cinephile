package api

import (
	"time"
)

type Thread struct {
	Thread_id      int       `json:"threadId"`
	Movie_id       int       `json:"movieId"`
	Parent_id      int       `json:"parentId,default=-1"`
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
	Channel_id      int   `json:"channelId"`
	Movie           Movie `json:"Movie"`
	Thread_count    int   `json:"threadCount"`
	Subscribe_count int   `json:"subscribeCount"`
	Like_count      int   `json:"likeCount"`
}
type Movie struct {
	Movie_id       int       `json:"movieId"`
	Channel_id     int       `json:"channelId"`
	Is_adult       bool      `json:"isAdult"`
	Original_title string    `json:"originalTitle"`
	Kr_title       string    `json:"krTitle"`
	Poster_path    string    `json:"posterPath"`
	Release_date   time.Time `json:"releaseDate"`
	Overview       string    `json:"overview"`
}
type Character struct {
	Name               string  `json:"name"`
	Role               string  `json:"role"`
	RepresentativeFilm []Movie `json:"representativeFilm"`
}
type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Image int    `json:"image"`
}
type RecommendForm struct {
	Thread_id int    `json:"threadId"`
	Email     string `json:"email"`
}
