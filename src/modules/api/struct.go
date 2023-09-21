package api

import (
	"time"
)

type Thread struct {
	Thread_id      int       `json:"threadId"`
	Parent_id      int       `json:"parentId,default=-1"`
	Title          string    `json:"title"`
	Content        string    `json:"content"`
	Created_at     time.Time `json:"createdAt"`
	Updated_at     time.Time `json:"updatedAt"`
	Like           int       `json:"like"`
	Is_recommended bool      `json:"isLiked"`
	Channel        Channel   `json:"channel"`
	Author         User      `json:"author"`
}
type ThreadRegistForm struct {
	Channel_id int    `json:"channelId"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Parent_id  int    `json:"parentId"`
	Is_exposed bool   `json:"isExposed"`
}
type Channel struct {
	Channel_id      int   `json:"channelId"`
	Thread_count    int   `json:"threadCount"`
	Subscribe_count int   `json:"subscribeCount"`
	Like_count      int   `json:"likeCount"`
	Movie           Movie `json:"movie"`
}
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
type Genre struct {
	Genre_id   int    `json:"genreId"`
	Genre_name string `json:"genreName"`
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
type Token struct {
	AccessToken   string `json:"access_token"`
	RefreshToken  string `json:"refresh_token"`
	Expire        int    `json:"expires_in"`
	RefreshExpire int    `json:"refresh_expires_in"`
	Type          string `json:"token_type"`
}
type KakaoTokenInfo struct {
	Id     string `json:"id"`
	Expire int    `json:"expires_in"`
	App_id int    `json:"app_id"`
}