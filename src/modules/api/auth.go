package api

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"os"

	. "cinephile/modules/dto"
	oauth "cinephile/modules/oauth"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/kakao"
)

const (
	CallBackURL = "http://localhost:4000/login/google/callback"

	// 인증 후 유저 정보를 가져오기 위한 API
	UserInfoAPIEndpoint = "https://www.googleapis.com/oauth2/v3/userinfo"

	// 인증 권한 범위. 여기에서는 프로필 정보 권한만 사용
	ScopeEmail   = "https://www.googleapis.com/auth/userinfo.email"
	ScopeProfile = "https://www.googleapis.com/auth/userinfo.profile"
)

var GoogleOAuthConf *oauth2.Config
var KakaoOAuthConf *oauth2.Config

func init() {
	if os.Getenv(`env`) == "" {
		godotenv.Load(`.env.local`)
	}
	GoogleOAuthConf = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_CALLBACK_URL"),
		Scopes:       []string{ScopeEmail, ScopeProfile},
		Endpoint:     google.Endpoint,
	}
	KakaoOAuthConf = &oauth2.Config{
		ClientID:     os.Getenv("KAKAO_CLIENT_ID"),
		ClientSecret: os.Getenv("KAKAO_CLIENT_SECRET"),
		Endpoint:     kakao.Endpoint,
		RedirectURL:  os.Getenv("KAKAO_CALLBACK_URL"),
	}
}

// state 값과 함께 Google 로그인 링크 생성
func GetLoginURL(state string) string {
	return GoogleOAuthConf.AuthCodeURL(state)
}

// 랜덤 state 생성기
func RandToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}
func GetGoogleInfo(c *gin.Context) (OauthInfo, error) {
	token, err := c.Cookie(`access_token`)
	if err != nil {
		return OauthInfo{}, err
	}
	return oauth.GetGoogleTokenInfo(token)
}
func GetKakaoInfo(c *gin.Context) (OauthInfo, error) {
	token, err := c.Cookie(`access_token`)
	if err != nil {
		return OauthInfo{}, err
	}
	return oauth.GetKakaoTokenInfo(token)
}
func getInfo(token string, platform string) (int, error) {
	if platform == "kakao" {
		return getKakaoTokenID(token)
	}
	if platform == "google" {
		return getGoogleTokenID(token)
	}
	return 0, errors.New("Invalid platform")
}
func getKakaoTokenID(token string) (int, error) {
	return oauth.GetKakaoTokenID(token)
}
func getGoogleTokenID(token string) (int, error) {
	return oauth.GetGoogleTokenID(token)
}
func getKakaoInfo(token string) (OauthInfo, error) {
	return oauth.GetKakaoTokenInfo(token)
}
func getGoogleInfo(token string) (OauthInfo, error) {
	return oauth.GetGoogleTokenInfo(token)
	// return OauthInfo{}, nil
}
func GetOAuthInfo(token string, platform string) (OauthInfo, error) {
	if platform == "kakao" {
		return getKakaoInfo(token)
	}
	if platform == "google" {
		return getGoogleInfo(token)
	}
	return OauthInfo{}, errors.New("Invalid platform")
}
func OAuthLogin(c *gin.Context) (Token, error) {
	platform, valid := c.GetQuery(`platform`)
	if !valid {
		return Token{}, errors.New("Invalid platform parameter")
	}
	code, valid := c.GetQuery(`code`)
	if !valid {
		return Token{}, errors.New("No auth code")
	}
	if platform == "kakao" {
		return oauth.KakaoLogin(code)
	}
	if platform == "google" {
		return oauth.GoogleLogin(code)
	}
	return Token{}, nil
}
