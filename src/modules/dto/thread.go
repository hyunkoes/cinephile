package dto

import "time"

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
type RecommendForm struct {
	Thread_id int    `json:"threadId"`
	Email     string `json:"email"`
}
