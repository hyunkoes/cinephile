package api

import (
	"database/sql"
	"errors"

	ErrChecker "cinephile/modules/errors"
	"cinephile/modules/storage"

	"github.com/gin-gonic/gin"
)

//	func GetThreads(c *gin.Context) ([]Thread, error) {
//		db := storage.DB()
//		var length int
//		_ = db.QueryRow(`select count(*) from thread`).Scan(&length)
//		if length == 0 {
//			return []Thread{}, errors.New("Nothing to show")
//		}
//		rows, err := db.Query(`select * from thread`)
//		if err := ErrChecker.Check(err); err != nil {
//			return []Thread{}, err
//		}
//		defer rows.Close()
//		Threads := make([]Thread, 0)
//		var thread Thread
//		for rows.Next() {
//			err = rows.Scan(&thread.Thread_id, &thread.Channel.Id, &thread.Content,
//				&thread.Author.Id, &thread.Parent, &thread.Created_at, &thread.Updated_at)
//			if err := ErrChecker.Check(err); err != nil {
//				return []Thread{}, err
//			}
//			Threads = append(Threads, thread)
//		}
//		return Threads, nil
//	}
func GetChildThreadsWithRecommend(c *gin.Context) ([]Thread, error, int) {
	db := storage.DB()
	cursor, valid := c.GetQuery("cursor")
	if cursor == "-1" || !valid {
		cursor = "2147483647"
	}
	parent_id, valid := c.GetQuery("parent_id")
	var length int
	_ = db.QueryRow(`select count(*) from thread where thread_id < ` + cursor + ` and parent = ` + parent_id).Scan(&length)
	if length == 0 {
		return []Thread{}, errors.New("Nothing to show"), 0
	}
	query := `
	SELECT
		t.thread_id,
		t.channel_id,
		m.original_title,
		m.kr_title,
		m.movie_id,
		m.poster_path,
		t.email,
		t.parent,
		t.content,
		tr.is_recommended,
		t.updated_at
	FROM
		thread AS t
	LEFT JOIN
		thread_recommend AS tr ON t.email = tr.email and t.thread_id = tr.thread_id
	LEFT JOIN
		channel AS c ON t.channel_id = c.channel_id
	LEFT JOIN
		movie AS m ON c.movie_id = m.movie_id
	WHERE
		t.parent = ` + parent_id + `  
		and 
		t.thread_id < ` + cursor + `
	ORDER BY
		t.thread_id DESC
	LIMIT 10;
	`
	rows, err := db.Query(query)

	if err := ErrChecker.Check(err); err != nil {
		return []Thread{}, err, 0
	}
	defer rows.Close()
	Threads := make([]Thread, 0)
	var thread Thread
	var is_recommended sql.NullBool
	for rows.Next() {
		err = rows.Scan(&thread.Thread_id, &thread.Channel.Channel_id, &thread.Channel.Movie.Original_title,
			&thread.Channel.Movie.Kr_title, &thread.Channel.Movie.Movie_id, &thread.Channel.Movie.Poster_path, &thread.Author.Id, &thread.Parent_id, &thread.Content, &is_recommended, &thread.Updated_at)
		if err := ErrChecker.Check(err); err != nil {
			return []Thread{}, err, 0
		}
		if !is_recommended.Valid {
			thread.Is_recommended = false
		} else {
			thread.Is_recommended = is_recommended.Bool
		}
		Threads = append(Threads, thread)
	}
	last_cursor := Threads[len(Threads)-1].Thread_id
	return Threads, nil, last_cursor
}
func GetThread(c *gin.Context) (Thread_detail, error) {
	thread_id, valid := c.GetQuery("thread_id")
	if !valid {
		return Thread_detail{}, errors.New("Invalid query string")
	}
	db := storage.DB()
	query := `
	SELECT
		t.thread_id,
		t.channel_id,
		m.original_title,
		m.kr_title,
		m.movie_id,
		t.email,
		t.parent,
		t.content,
		tr.is_recommended,
		t.updated_at
	FROM
		thread AS t
	LEFT JOIN
		thread_recommend AS tr ON t.email = tr.email and t.thread_id = tr.thread_id
	LEFT JOIN
		channel AS c ON t.channel_id = c.channel_id
	LEFT JOIN
		movie AS m ON c.movie_id = m.movie_id
	WHERE
		t.thread_id = ` + thread_id + `;
	`

	child_query := `
	SELECT
		t.thread_id,
		t.channel_id,
		m.original_title,
		m.kr_title,
		m.movie_id,
		t.email,
		t.parent,
		t.content,
		tr.is_recommended,
		t.updated_at
	FROM
		thread AS t
	LEFT JOIN
		thread_recommend AS tr ON t.email = tr.email and t.thread_id = tr.thread_id
	LEFT JOIN
		channel AS c ON t.channel_id = c.channel_id
	LEFT JOIN
		movie AS m ON c.movie_id = m.movie_id
	WHERE
		t.parent = ` + thread_id + `
	ORDER BY
		t.thread_id;
	`
	var thread Thread_detail
	var is_recommended sql.NullBool
	err := db.QueryRow(query).Scan(&thread.Self.Thread_id, &thread.Self.Channel.Channel_id, &thread.Self.Channel.Movie.Original_title,
		&thread.Self.Channel.Movie.Kr_title, &thread.Self.Channel.Movie.Movie_id, &thread.Self.Author.Id, &thread.Self.Parent_id, &thread.Self.Content, &is_recommended, &thread.Self.Updated_at)
	if !is_recommended.Valid {
		thread.Self.Is_recommended = false
	} else {
		thread.Self.Is_recommended = is_recommended.Bool
	}
	rows, err := db.Query(child_query)
	if err := ErrChecker.Check(err); err != nil {
		return Thread_detail{}, err
	}
	parent_id := thread.Self.Parent_id
	parent_query := `
	SELECT
		t.thread_id,
		t.channel_id,
		m.original_title,
		m.kr_title,
		m.movie_id,
		t.email,
		t.parent,
		t.content,
		tr.is_recommended,
		t.updated_at
	FROM
		thread AS t
	LEFT JOIN
		thread_recommend AS tr ON t.email = tr.email and t.thread_id = tr.thread_id
	LEFT JOIN
		channel AS c ON t.channel_id = c.channel_id
	LEFT JOIN
		movie AS m ON c.movie_id = m.movie_id
	WHERE
		t.thread_id = ` + string(parent_id) + `;`
	if parent_id != -1 {
		err = db.QueryRow(parent_query).Scan(&thread.Parent.Thread_id, &thread.Parent.Channel.Channel_id, &thread.Parent.Channel.Movie.Original_title,
			&thread.Parent.Channel.Movie.Kr_title, &thread.Parent.Channel.Movie.Movie_id, &thread.Parent.Author.Id, &thread.Parent.Parent_id, &thread.Parent.Content, &is_recommended, &thread.Parent.Updated_at)
		if !is_recommended.Valid {
			thread.Parent.Is_recommended = false
		} else {
			thread.Parent.Is_recommended = is_recommended.Bool
		}
	}

	var child_thread Thread
	thread.Child = make([]Thread, 0)

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&child_thread.Thread_id, &child_thread.Channel.Channel_id, &child_thread.Channel.Movie.Original_title, &child_thread.Channel.Movie.Kr_title, &child_thread.Channel.Movie.Movie_id,
			&child_thread.Author.Id, &child_thread.Parent_id, &child_thread.Content, &is_recommended, &child_thread.Updated_at)
		if err := ErrChecker.Check(err); err != nil {
			return Thread_detail{}, err
		}
		if !is_recommended.Valid {
			child_thread.Is_recommended = false
		} else {
			child_thread.Is_recommended = is_recommended.Bool
		}
		thread.Child = append(thread.Child, child_thread)
	}

	return thread, nil
}

func GetThreadsWithRecommend(c *gin.Context) ([]Thread, error, int) {
	db := storage.DB()
	cursor, valid := c.GetQuery("cursor")
	if cursor == "-1" || !valid {
		cursor = "2147483647"
	}
	var length int
	_ = db.QueryRow(`select count(*) from thread where thread_id < ` + cursor).Scan(&length)
	if length == 0 {
		return []Thread{}, errors.New("Nothing to show"), 0
	}
	query := `
	SELECT
		t.thread_id,
		t.channel_id,
		m.original_title,
		m.kr_title,
		m.movie_id,
		m.poster_path,
		t.email,
		t.parent,
		t.content,
		tr.is_recommended,
		t.updated_at
	FROM
		thread AS t
	LEFT JOIN
		thread_recommend AS tr ON t.email = tr.email and t.thread_id = tr.thread_id
	LEFT JOIN
		channel AS c ON t.channel_id = c.channel_id
	LEFT JOIN
		movie AS m ON c.movie_id = m.movie_id
	WHERE t.thread_id < ` + cursor + `
	ORDER BY
		t.thread_id DESC
	LIMIT 10;
	`
	rows, err := db.Query(query)

	if err := ErrChecker.Check(err); err != nil {
		return []Thread{}, err, 0
	}
	defer rows.Close()
	Threads := make([]Thread, 0)
	var thread Thread
	var is_recommended sql.NullBool
	for rows.Next() {
		err = rows.Scan(&thread.Thread_id, &thread.Channel.Channel_id, &thread.Channel.Movie.Original_title,
			&thread.Channel.Movie.Kr_title, &thread.Channel.Movie.Movie_id, &thread.Channel.Movie.Poster_path, &thread.Author.Id, &thread.Parent_id, &thread.Content, &is_recommended, &thread.Updated_at)
		if err := ErrChecker.Check(err); err != nil {
			return []Thread{}, err, 0
		}
		if !is_recommended.Valid {
			thread.Is_recommended = false
		} else {
			thread.Is_recommended = is_recommended.Bool
		}
		Threads = append(Threads, thread)
	}
	last_cursor := Threads[len(Threads)-1].Thread_id
	return Threads, nil, last_cursor
}

func RegistThread(c *gin.Context) error {
	var reqBody Thread
	user := c.GetHeader("user")
	err := c.ShouldBind(&reqBody)

	if ErrChecker.Check(err) != nil {
		return err
	}
	if reqBody.Parent_id == 0 {
		reqBody.Parent_id = -1
	}
	db := storage.DB()
	_, err = db.Exec(`Insert into thread (channel_id,content,email,parent) values(?,?,?,?)`, reqBody.Channel.Channel_id, reqBody.Content, user, reqBody.Parent_id)
	if err := ErrChecker.Check(err); err != nil {
		return err
	}

	return nil
}

func ChangeRecommendThread(c *gin.Context) error {
	var reqBody RecommendForm
	err := c.ShouldBind(&reqBody)

	if ErrChecker.Check(err) != nil {
		return err
	}
	db := storage.DB()
	var is_recommended bool
	err = db.QueryRow(`select is_recommended from thread_recommend where thread_id = ? and email = ? `, reqBody.Thread_id, reqBody.Email).Scan(&is_recommended)

	if err == sql.ErrNoRows {
		// No row -> is_recommended : true
		_, err = db.Exec(`Insert into thread_recommend (thread_id, email, is_recommended) values(?,?,true)`, reqBody.Thread_id, reqBody.Email)
		if err := ErrChecker.Check(err); err != nil {
			return err
		}
		// row 추가
		return nil
	}
	// Is_recommended -> is_recommended : false
	if is_recommended == true {
		_, err = db.Exec(`Update thread_recommend set is_recommended = false where thread_id = ? and email = ?`, reqBody.Thread_id, reqBody.Email)
		if err := ErrChecker.Check(err); err != nil {
			return err
		}
		return nil
	}
	// Not is_recommneded -> is_recommend : true
	_, err = db.Exec(`Update thread_recommend set is_recommended = true where thread_id = ? and email = ?`, reqBody.Thread_id, reqBody.Email)
	if err := ErrChecker.Check(err); err != nil {
		return err
	}
	return nil
}

func DeleteThread(c *gin.Context) error {
	thread_id := c.GetHeader("thread_id")
	user := c.GetHeader("user")
	query := `
	SELECT
		count(*)
	FROM 
		thread
	WHERE
		thread_id = ` + thread_id + `
		and 
		email = "` + user + `"
	`
	db := storage.DB()
	var length int
	_ = db.QueryRow(query).Scan(&length)
	if length == 0 {
		return errors.New("Incorrect user and thread")
	}
	query = `
	DELETE FROM
		thread
	WHERE 
		thread_id = ` + thread_id + `
		and 
		email = "` + user + `"
	`
	_, err := db.Exec(query)
	if err := ErrChecker.Check(err); err != nil {
		return err
	}
	return nil
}
