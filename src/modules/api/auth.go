package api

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
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
func getKakaoInfo(token string) (string, error) {
	header := "Bearer " + token
	_ = header

	return "", nil
}
func OAuthLogin(c *gin.Context) (Token, error) {
	code, valid := c.GetQuery(`code`)
	if !valid {
		return Token{}, errors.New("No auth code")
	}

	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("client_id", KakaoOAuthConf.ClientID)
	data.Set("client_secret", KakaoOAuthConf.ClientSecret)
	data.Set("redirect_uri", KakaoOAuthConf.RedirectURL)
	data.Set("code", code) // 사용자 인증 후 받은 인증 코드
	platform := c.GetHeader(`platform`)
	_ = platform
	// POST 요청 보내기
	resp, err := http.PostForm(KakaoOAuthConf.Endpoint.TokenURL, data)
	if err != nil {
		fmt.Printf("Kakao token API err : %v\n", err)
	}
	defer resp.Body.Close()

	// 응답 바디 읽기
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Kakao token response body err : %v\n", err)
	}
	var token Token
	if err := json.Unmarshal(body, &token); err != nil {
		return Token{}, err
	}
	return token, nil
}
