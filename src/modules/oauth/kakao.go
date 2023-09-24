package oauth

import (
	. "cinephile/const"
	. "cinephile/modules/dto"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/kakao"
)

var KakaoOAuthConf *oauth2.Config

func init() {
	if os.Getenv(`env`) == "" {
		godotenv.Load(`.env.local`)
	}
	KakaoOAuthConf = &oauth2.Config{
		ClientID:     os.Getenv("KAKAO_CLIENT_ID"),
		ClientSecret: os.Getenv("KAKAO_CLIENT_SECRET"),
		Endpoint:     kakao.Endpoint,
		RedirectURL:  os.Getenv("KAKAO_CALLBACK_URL"),
	}
}
func GetKakaoTokenID(token string) (int, error) {
	apiURL := KAKAO_GET_TOKEN_ID_URL
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return -1, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return -1, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1, err
	}
	var kakaoTokenInfo KakaoTokenInfo
	if err := json.Unmarshal(body, &kakaoTokenInfo); err != nil {
		return -1, err
	}
	return kakaoTokenInfo.Id, nil
}
func GetKakaoTokenInfo(token string) (OauthInfo, error) {
	var kakaoInfo OauthInfo
	apiURL := KAKAO_GET_INFO_URL
	req, err := http.NewRequest("POST", apiURL, nil)
	if err != nil {
		return OauthInfo{}, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	resp, err := http.DefaultClient.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return OauthInfo{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return OauthInfo{}, err
	}
	var payload interface{}                      //The interface where we will save the converted JSON data.
	_ = json.Unmarshal(body, &payload)           // Convert JSON data into interface{} type
	jsonData := payload.(map[string]interface{}) // To use the converted data we will need to convert it into a map[string]interface
	kakaoID := jsonData["id"].(float64)          // id는 숫자로 반환되기 때문에 float64로 형변환
	profile := jsonData["kakao_account"].(map[string]interface{})["profile"].(map[string]interface{})
	name := profile["nickname"].(string)
	photo := profile["thumbnail_image_url"].(string)

	kakaoInfo.ID = strconv.Itoa(int(kakaoID))
	kakaoInfo.Name = name
	kakaoInfo.Image = photo
	return kakaoInfo, nil
}
func KakaoLogin(code string) (Token, error) {
	data := url.Values{}
	data.Set("Content-Type", "application/x-www-form-urlencoded")
	data.Set("grant_type", "authorization_code")
	data.Set("client_id", KakaoOAuthConf.ClientID)
	data.Set("client_secret", KakaoOAuthConf.ClientSecret)
	data.Set("redirect_uri", KakaoOAuthConf.RedirectURL)
	data.Set("code", code) // 사용자 인증 후 받은 인증 코드
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
func KakaoRefreshATwithRT(refresh_token string) (Token, error) {
	apiUrl := KAKAO_TOKEN_REFRESH_URL
	_ = apiUrl
	data := url.Values{}
	data.Set("Content-Type", "application/x-www-form-urlencoded")
	data.Set("grant_type", "refresh_token")
	data.Set("client_id", KakaoOAuthConf.ClientID)
	data.Set("refresh_token", refresh_token)
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
