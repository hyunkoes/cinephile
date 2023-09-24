package dto

type User struct {
	Id       string `json:"id"`
	Platform string `json:"platform"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Image    string `json:"image"`
}
