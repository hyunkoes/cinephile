package middleware

import (
	"cinephile/modules/middleware/token"

	"github.com/gin-gonic/gin"
)

func TokenCheck(c *gin.Context) error {
	if token.AccessTokenIsValid(c) {
		return nil
	}
	if token.RefreshTokenIsValid(c) {
		refresh_token, err := c.Cookie(`refresh_token`)
		if err != nil {
			return err
		}
		tokens, err := token.RefreshAT(refresh_token)
		if err != nil {
			return err
		}
		platform, err := c.Cookie(`platform`)
		if err != nil {
			return err
		}
		c.SetCookie("access_token", tokens.AccessToken, tokens.Expire, "/", "", false, true)
		c.SetCookie("refresh_token", tokens.RefreshToken, tokens.RefreshExpire, "/", "", false, true)
		c.SetCookie("platform", platform, tokens.RefreshExpire, "/", "", false, true)
	}
	return nil
}
