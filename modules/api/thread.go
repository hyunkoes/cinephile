package api

import (
	"errors"
	"fmt"

	ErrChecker "cinephile/modules/errors"
	"cinephile/modules/storage"

	"github.com/gin-gonic/gin"
)

func GetThreads(c *gin.Context) ([]Thread, error) {
	db := storage.DB()
	var length int
	_ = db.QueryRow(`select count(*) from thread`).Scan(&length)
	if length == 0 {
		return []Thread{}, errors.New("Nothing to show")
	}
	rows, err := db.Query(`select * from thread`)
	if err := ErrChecker.Check(err); err != nil {
		return []Thread{}, err
	}
	defer rows.Close()
	Threads := make([]Thread, 0)
	var thread Thread
	for rows.Next() {
		err = rows.Scan(&thread.Thread_id, &thread.Channel_id, &thread.Content,
			&thread.Email, &thread.Parent, &thread.Created_at, &thread.Updated_at)
		if err := ErrChecker.Check(err); err != nil {
			return []Thread{}, err
		}
		Threads = append(Threads, thread)
	}
	return Threads, nil
}

func GetThreadsWithRecommend(c *gin.Context) ([]Thread_recommend, error) {
	db := storage.DB()
	var length int
	_ = db.QueryRow(`select count(*) from thread`).Scan(&length)
	if length == 0 {
		return []Thread_recommend{}, errors.New("Nothing to show")
	}
	rows, err := db.Query(`select * from thread left join thread_recommend on thread.email = thread_recommend.email`)
	if err := ErrChecker.Check(err); err != nil {
		return []Thread_recommend{}, err
	}
	defer rows.Close()
	Threads := make([]Thread_recommend, 0)
	var thread Thread_recommend
	for rows.Next() {
		err = rows.Scan(&thread.Thread_id, &thread.Channel_id, &thread.Content,
			&thread.Email, &thread.Parent, &thread.Created_at, &thread.Updated_at, &thread.Is_recommend)
		if err := ErrChecker.Check(err); err != nil {
			return []Thread_recommend{}, err
		}
		Threads = append(Threads, thread)
	}
	return Threads, nil
}

func RegistThread(c *gin.Context) error {
	var reqBody Thread
	err := c.ShouldBind(&reqBody)

	if ErrChecker.Check(err) != nil {
		return err
	}
	db := storage.DB()
	_, err = db.Exec(`Insert into thread (channel_id,content,email) values(?,?,?)`, reqBody.Channel_id, reqBody.Content, reqBody.Email)
	if err := ErrChecker.Check(err); err != nil {
		return err
	}
	fmt.Println(reqBody.Content)

	return nil
}
