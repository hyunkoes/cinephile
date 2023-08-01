package api

import (
	"errors"

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
			&thread.Email, &thread.Created_at, &thread.Updated_at)
		if err := ErrChecker.Check(err); err != nil {
			Threads = append(Threads, thread)
		}
		return []Thread{}, err
	}
	return Threads, nil
}
