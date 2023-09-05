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
		Name     string `json:"name`
	}
	err := c.ShouldBind(&form)
	if err != nil {
		return err
	}
	name := form.Email
	if form.Name != "" {
		name = form.Name
	}
	query := fmt.Sprintf(`insert into user values ("%s", "%s", "%s", %d)`, form.Email, form.Password, name, 1)
	db := storage.DB()
	_, err = db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
