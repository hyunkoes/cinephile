package api

import (
	"cinephile/modules/storage"
	"fmt"

	"github.com/gin-gonic/gin"
)

func RegistUserForTest(c *gin.Context) error {
	var form struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := c.ShouldBind(&form)
	if err != nil {
		return err
	}
	query := fmt.Sprintf(`insert into user values ("%s", "%s", "%s", %d)`, form.Email, form.Password, form.Email, 1)
	db := storage.DB()
	_, err = db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
