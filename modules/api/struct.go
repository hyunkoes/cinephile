package api

import "time"

type Thread struct {
	Thread_id  int       `json:"thread_id"`
	Channel_id int       `json:"channel_id"`
	Content    string    `json:"content"`
	Email      string    `json:"email"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
