package oauth

func GetID(token string, platform string) (int, error) {
	if platform == "kakao" {
		return GetKakaoTokenID(token)
	}
	if platform == "google" {
		return GetGoogleTokenID(token)
	}
	return 0, nil
}
