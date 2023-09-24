package api

import (
	. "cinephile/modules/dto"
	"cinephile/modules/logging"
	"cinephile/modules/storage"
	"database/sql"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegistUser(OauthInfo OauthInfo, platform string) error {
	db := storage.DB()
	_, err := db.Exec(`Insert into user (id, platform, user_name, photo) values(?,?,?,?)`, OauthInfo.ID, platform, OauthInfo.Name, OauthInfo.Image)
	if err != nil {
		logging.Warn(err)
		return err
	}
	return nil
}

func GetUser(c *gin.Context) (User, error) {
	token, err := c.Cookie(`access_token`)
	if err != nil {
		return User{}, err
	}
	platform, valid := c.GetQuery(`platform`)
	if !valid {
		return User{}, errors.New("Input platform(kakao, google ..) in parameter !")
	}
	id, err := getInfo(token, platform)
	if err != nil {
		return User{}, err
	}
	wrapId := strconv.Itoa(id)
	db := storage.DB()
	query := `
		select * from user where id = "` + wrapId + `" and platform = "` + platform + `"`
	row := db.QueryRow(query)
	var user User
	var password sql.NullString
	err = row.Scan(&user.Id, &user.Platform, &password, &user.Name, &user.Image)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func IsExistUser(oauthInfo OauthInfo, platform string) bool {
	db := storage.DB()
	var length int
	_ = db.QueryRow(`select count(*) from user where id = "?" and platform ="?"`, oauthInfo.ID, platform).Scan(length)
	return length > 0
}

func GetUsers() ([]User, error) {
	users := make([]User, 0)
	db := storage.DB()
	rows, err := db.Query(`select * from user`)
	if err != nil {
		return []User{}, err
	}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Password, &user.Name, &user.Image)
		if err != nil {
			return []User{}, err
		}
		users = append(users, user)
	}
	return users, nil
}
