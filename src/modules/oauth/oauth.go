package oauth

func GetID(token string, platform string) (string, error) {
	if platform == "kakao" {
		return GetKakaoTokenID(token)
	}
	if platform == "google" {
		return GetGoogleTokenID(token)
	}
	return "", nil
}
