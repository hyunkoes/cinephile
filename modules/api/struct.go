package api

import "time"

type Thread struct {
	Thread_id  int       `json:"thread_id"`
	Channel_id int       `json:"channel_id"`
	Content    string    `json:"content"`
	Email      string    `json:"email"`
	Parent     int       `json:"parent"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
type Thread_recommend struct {
	Thread_id    int       `json:"thread_id"`
	Channel_id   int       `json:"channel_id"`
	Content      string    `json:"content"`
	Email        string    `json:"email"`
	Parent       int       `json:"parent"`
	Created_at   time.Time `json:"created_at"`
	Updated_at   time.Time `json:"updated_at"`
	Is_recommend bool      `json:"is_recommend"`
}
type RecommendForm struct {
	Thread_id int    `json:"thread_id"`
	Email     string `json:"email"`
}
