package token

import (
	. "cinephile/modules/dto"
	"cinephile/modules/oauth"
)

func RefreshAT(refresh_token string, platform string) (Token, error) {
	var token Token
	var err error
	if platform == "kakao" {
		token, err = oauth.KakaoRefreshATwithRT(refresh_token)
	}
	if platform == "google" {
		token, err = oauth.GoogleRefreshATwithRT(refresh_token)
	}
	if err != nil {
		return Token{}, err
	}
	return token, nil
}
