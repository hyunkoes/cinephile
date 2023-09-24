package token

import (
	. "cinephile/modules/dto"
	"cinephile/modules/oauth"
)

func RefreshAT(refresh_token string) (Token, error) {
	token, err := oauth.KakaoRefreshATwithRT(refresh_token)
	if err != nil {
		return Token{}, err
	}
	return token, nil
}
