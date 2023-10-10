package token

import "github.com/gin-gonic/gin"

func AccessTokenIsValid(c *gin.Context) bool {
	_, err := c.Cookie(`access_token`)
	return err == nil
}

func RefreshTokenIsValid(c *gin.Context) bool {
	_, err := c.Cookie(`refresh_token`)
	return err == nil
}
