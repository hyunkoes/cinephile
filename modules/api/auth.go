package api

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	CallBackURL = "http://localhost:4000/login/google/callback"

	// 인증 후 유저 정보를 가져오기 위한 API
	UserInfoAPIEndpoint = "https://www.googleapis.com/oauth2/v3/userinfo"

	// 인증 권한 범위. 여기에서는 프로필 정보 권한만 사용
	ScopeEmail   = "https://www.googleapis.com/auth/userinfo.email"
	ScopeProfile = "https://www.googleapis.com/auth/userinfo.profile"
)

var OAuthConf *oauth2.Config

func init() {
	OAuthConf = &oauth2.Config{
		ClientID:     "902520195219-bo2hk1ap7sdholam80b930dgglebocnu.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL:  CallBackURL,
		Scopes:       []string{ScopeEmail, ScopeProfile},
		Endpoint:     google.Endpoint,
	}
}

// state 값과 함께 Google 로그인 링크 생성
func GetLoginURL(state string) string {
	return OAuthConf.AuthCodeURL(state)
}

// 랜덤 state 생성기
func RandToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}
