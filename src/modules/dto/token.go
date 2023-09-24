package dto

type Token struct {
	AccessToken   string `json:"access_token"`
	RefreshToken  string `json:"refresh_token"`
	Expire        int    `json:"expires_in"`
	RefreshExpire int    `json:"refresh_expires_in"`
	Type          string `json:"token_type"`
}
type KakaoTokenInfo struct {
	Id     int `json:"id"`
	Expire int `json:"expires_in"`
	App_id int `json:"app_id"`
}
type OauthInfo struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}
