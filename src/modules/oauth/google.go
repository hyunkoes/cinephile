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

	"golang.org/x/oauth2/google"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

var GoogleOAuthConf *oauth2.Config

func init() {
	if os.Getenv(`env`) == "" {
		godotenv.Load(`.env.local`)
	}
	GoogleOAuthConf = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Endpoint:     google.Endpoint,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		RedirectURL: os.Getenv("GOOGLE_CALLBACK_URL"),
	}
}
func GetGoogleTokenInfo(token string) (OauthInfo, error) {
	var googleInfo OauthInfo
	apiURL := GOOGLE_GET_INFO_URL
	req, err := http.NewRequest("GET", apiURL, nil)
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
	fmt.Println(string(body))
	var payload interface{}                      //The interface where we will save the converted JSON data.
	_ = json.Unmarshal(body, &payload)           // Convert JSON data into interface{} type
	jsonData := payload.(map[string]interface{}) // To use the converted data we will need to convert it into a map[string]interface
	googleID := jsonData["id"].(float64)         // id는 숫자로 반환되기 때문에 float64로 형변환
	name := jsonData["name"].(string)
	photo := jsonData["picture"].(string)

	googleInfo.ID = strconv.Itoa(int(googleID))
	googleInfo.Name = name
	googleInfo.Image = photo
	return googleInfo, nil
}

func GetGoogleTokenID(token string) (int, error) {
	return -1, nil
}

func GoogleRefreshATwithRT(token string) (Token, error) {
	return Token{}, nil
}
func GoogleLogin(code string) (Token, error) {
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("client_id", GoogleOAuthConf.ClientID)
	data.Set("client_secret", GoogleOAuthConf.ClientSecret)
	data.Set("redirect_uri", GoogleOAuthConf.RedirectURL)
	data.Set("code", code) // 사용자 인증 후 받은 인증 코드
	// POST 요청 보내기
	resp, err := http.PostForm(GoogleOAuthConf.Endpoint.TokenURL, data)
	fmt.Println("POST 요청 보냄!")
	if err != nil {
		fmt.Printf("Google token API err : %v\n", err)
	}
	defer resp.Body.Close()

	// 응답 바디 읽기
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("POST 요청 받음!")
	fmt.Println(string(body))
	if err != nil {
		fmt.Printf("Google token response body err : %v\n", err)
	}
	var token Token
	if err := json.Unmarshal(body, &token); err != nil {
		return Token{}, err
	}
	return token, nil
}
