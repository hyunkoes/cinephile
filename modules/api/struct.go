package api

import (
	"time"
)

type Thread struct {
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

type Thread_detail struct {
	Self   Thread   `json:"self"`
	Parent Thread   `json:"parent"`
	Child  []Thread `json:"child"`
}

type Thread_list struct {
	List []Thread `json:"list"`
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

type Channel struct {
}
