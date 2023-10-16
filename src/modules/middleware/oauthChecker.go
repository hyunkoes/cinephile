package middleware

import (
	. "cinephile/const"
	"cinephile/modules/middleware/token"
	"cinephile/modules/oauth"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func TokenCheck(c *gin.Context) {
	if token.AccessTokenIsValid(c) {
		token, _ := c.Cookie(`access_token`)
		platform, err := c.Cookie(`platform`)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		id, err := oauth.GetID(token, platform)
		c.Set(`user`, id)
		c.Next()
		return
	}
	if token.RefreshTokenIsValid(c) {
		// Get Refresh token
		refresh_token, err := c.Cookie(`refresh_token`)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		// Get platform
		platform, err := c.Cookie(`platform`)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		fmt.Println(refresh_token, platform, "으로 액세스 토큰 재발급 시작")
		// Refresh tokens by platform oauth server
		tokens, err := token.RefreshAT(refresh_token, platform)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		// If refresh token is refreshed, set again
		c.SetSameSite(http.SameSiteNoneMode)

		if tokens.RefreshToken != "" {
			c.SetCookie("refresh_token", tokens.RefreshToken, tokens.RefreshExpire, "/", COOKIE_DOMAIN, true, true)
			c.SetCookie("platform", platform, tokens.RefreshExpire, "/", COOKIE_DOMAIN, true, true)
		}
		// Set cookie for next handler ( request )
		c.Request.AddCookie(&http.Cookie{
			Name:     "access_token",
			Value:    tokens.AccessToken,
			Path:     "/",
			Domain:   COOKIE_DOMAIN,
			Expires:  time.Now().Add(time.Duration(tokens.Expire) * time.Second),
			MaxAge:   0,
			Secure:   true,
			HttpOnly: true,
		})
		// Set cookie for client ( response )
		c.SetCookie("access_token", tokens.AccessToken, tokens.Expire, "/", COOKIE_DOMAIN, true, true)
		user_id, err := oauth.GetID(tokens.AccessToken, platform)
		c.SetCookie("TEST1", "TEST", 10000, "/", COOKIE_DOMAIN, true, true)
		c.Set(`user`, user_id)
		c.Next()
		return
	}
	c.JSON(http.StatusUnauthorized, gin.H{"error": "No token. Login again !"})
	c.Abort()
	return
}
