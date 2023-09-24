package dto

type Channel struct {
	Channel_id      int   `json:"channelId"`
	Thread_count    int   `json:"threadCount"`
	Subscribe_count int   `json:"subscribeCount"`
	Like_count      int   `json:"likeCount"`
	Movie           Movie `json:"movie"`
}
