package middleware

import (
	"cinephile/modules/middleware/token"
	"cinephile/modules/oauth"
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
		c.Header(`user`, id)
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
		// Refresh tokens by platform oauth server
		tokens, err := token.RefreshAT(refresh_token, platform)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		// If refresh token is refreshed, set again
		if tokens.RefreshToken != "" {
			c.SetCookie("refresh_token", tokens.RefreshToken, tokens.RefreshExpire, "/", "", false, true)
			c.SetCookie("platform", platform, tokens.RefreshExpire, "/", "", false, true)
		}
		// Set cookie for next handler ( request )
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
		// Set cookie for client ( response )
		c.SetCookie("access_token", tokens.AccessToken, tokens.Expire, "/", "", false, true)
		user_id, err := oauth.GetID(tokens.AccessToken, platform)
		c.Header(`user`, user_id)
		c.Next()
		return
	}
	c.JSON(http.StatusUnauthorized, gin.H{"error": "No token. Login again !"})
	c.Abort()
	return
}
