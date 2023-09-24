package middleware

import (
	"cinephile/modules/middleware/token"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func TokenCheck(c *gin.Context) {
	if token.AccessTokenIsValid(c) {
		c.Next()
	}
	if token.RefreshTokenIsValid(c) {
		refresh_token, err := c.Cookie(`refresh_token`)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		tokens, err := token.RefreshAT(refresh_token)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		platform, err := c.Cookie(`platform`)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		// kakao는 refresh token의 유효기간이 1달보다 짧으면 refresh token이 재발급됨
		if tokens.RefreshToken != "" {
			c.SetCookie("refresh_token", tokens.RefreshToken, tokens.RefreshExpire, "/", "", false, true)
			c.SetCookie("platform", platform, tokens.RefreshExpire, "/", "", false, true)
		}
		// Next 핸들러에서 쿠키를 바로 사용하기 위함
		c.Request.AddCookie(&http.Cookie{
			Name:     "access_token",
			Value:    tokens.AccessToken,
			Path:     "/",
			Domain:   "",
			Expires:  time.Now().Add(time.Duration(tokens.Expire) * time.Second),
			MaxAge:   0,
			Secure:   false,
			HttpOnly: true,
		})
		// Response에 쿠키를 넘겨주어 Client에서 등록할 수 있도록 함
		c.SetCookie("access_token", tokens.AccessToken, tokens.Expire, "/", "", false, true)
		c.Next()
		return
	}
	c.JSON(http.StatusUnauthorized, gin.H{"error": "No token. Login again !"})
	c.Abort()
	return
}
