package api

import (
	"database/sql"
	"errors"

	ErrChecker "cinephile/modules/errors"
	"cinephile/modules/storage"
	. "cinephile/modules/tmdb"

	"github.com/gin-gonic/gin"
)

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
		return []Thread{}, nil, 0
	}
	query := `
	SELECT
		t.thread_id,
		t.channel_id,
		m.original_title,
		m.kr_title,
		m.movie_id,
		m.poster_path,
		u.email,
		u.user_name,
		u.photo,
		t.parent,
		t.content,
		t.like_count,
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
	LEFT JOIN
		user AS u ON u.email = t.email
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
			&thread.Channel.Movie.Kr_title, &thread.Channel.Movie.Movie_id, &thread.Channel.Movie.Poster_path, &thread.Author.Id, &thread.Author.Name, &thread.Author.Image, &thread.Parent_id, &thread.Content, &thread.Like, &is_recommended, &thread.Updated_at)
		if err := ErrChecker.Check(err); err != nil {
			return []Thread{}, err, 0
		}
		if !is_recommended.Valid {
			thread.Is_recommended = false
		} else {
			thread.Is_recommended = is_recommended.Bool
		}
		thread.Channel.Movie.Poster_path = TmdbPosterAPI(thread.Channel.Movie.Poster_path)
		Threads = append(Threads, thread)
	}
	last_cursor := Threads[len(Threads)-1].Thread_id
	return Threads, nil, last_cursor
}
func GetThread(c *gin.Context) (Thread, error) {
	thread_id, valid := c.GetQuery("thread_id")
	if !valid {
		return Thread{}, errors.New("Invalid query string")
	}
	db := storage.DB()
	query := `
	SELECT
		t.thread_id,
		t.channel_id,
		m.original_title,
		m.kr_title,
		m.movie_id,
		u.email,
		u.user_name,
		u.photo,
		t.parent,
		t.content,
		t.like_count,
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
	LEFT JOIN
		user AS u ON u.email = t.email
	WHERE
		t.thread_id = ` + thread_id + `;
	`
	var thread Thread
	var is_recommended sql.NullBool
	err := db.QueryRow(query).Scan(&thread.Thread_id, &thread.Channel.Channel_id, &thread.Channel.Movie.Original_title,
		&thread.Channel.Movie.Kr_title, &thread.Channel.Movie.Movie_id, &thread.Author.Id, &thread.Author.Name, &thread.Author.Image, &thread.Parent_id, &thread.Content, &thread.Like, &is_recommended, &thread.Updated_at)
	if !is_recommended.Valid {
		thread.Is_recommended = false
	} else {
		thread.Is_recommended = is_recommended.Bool
	}
	if err != nil {
		return Thread{}, err
	}
	return thread, nil
}

func GetThreadsWithRecommend(c *gin.Context) ([]Thread, error, int) {
	db := storage.DB()
	cursor, valid := c.GetQuery("cursor")
	if !valid {
		cursor = "2147483647"
	}
	var length int
	query := `select count(*) from thread where thread_id < ` + cursor
	_ = db.QueryRow(query).Scan(&length)
	if length == 0 {
		return []Thread{}, nil, 0
	}
	query = `
SELECT
	t.thread_id,
	t.channel_id,
	m.original_title,
	m.kr_title,
	m.movie_id,
	m.poster_path,
	u.email,
	u.user_name,
	u.photo,
	t.parent,
	t.content,
	t.like_count,
	tr.is_recommended,
	t.created_at,
	t.updated_at
FROM
	thread AS t
LEFT JOIN
	thread_recommend AS tr ON t.email = tr.email and t.thread_id = tr.thread_id
LEFT JOIN
	channel AS c ON t.channel_id = c.channel_id
LEFT JOIN
	movie AS m ON c.movie_id = m.movie_id
LEFT JOIN
	user AS u ON u.email = t.email
WHERE 
	t.is_exposed = true
	and
	t.thread_id < ` + cursor + `
ORDER BY
	t.thread_id DESC
LIMIT 11;
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
			&thread.Channel.Movie.Kr_title, &thread.Channel.Movie.Movie_id, &thread.Channel.Movie.Poster_path, &thread.Author.Id, &thread.Author.Name, &thread.Author.Image, &thread.Parent_id, &thread.Content, &thread.Like, &is_recommended, &thread.Created_at, &thread.Updated_at)
		if err := ErrChecker.Check(err); err != nil {
			return []Thread{}, err, 0
		}
		if !is_recommended.Valid {
			thread.Is_recommended = false
		} else {
			thread.Is_recommended = is_recommended.Bool
		}
		thread.Channel.Movie.Poster_path = TmdbPosterAPI(thread.Channel.Movie.Poster_path)
		Threads = append(Threads, thread)
	}
	last_cursor := -1
	if len(Threads) > 10 {
		Threads = Threads[:len(Threads)-1]
		last_cursor = Threads[len(Threads)-1].Thread_id
	}
	return Threads, nil, last_cursor
}

func RegistThread(c *gin.Context) error {
	var reqBody ThreadRegistForm
	user := c.GetHeader("user")
	err := c.ShouldBind(&reqBody)

	if ErrChecker.Check(err) != nil {
		return err
	}
	if reqBody.Parent_id == 0 {
		reqBody.Parent_id = -1
	}
	db := storage.DB()
	_, err = db.Exec(`Insert into thread (channel_id,title,content,email,parent,is_exposed) values(?,?,?,?,?,?)`, reqBody.Channel_id, reqBody.Title, reqBody.Content, user, reqBody.Parent_id, reqBody.Is_exposed)
	if err := ErrChecker.Check(err); err != nil {
		return err
	}
	_, err = db.Exec(`Update channel set thread_count = thread_count + 1 where channel_id = ?`, reqBody.Channel_id)
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
		_, err = db.Exec(`Update thread set like_count = like_count + 1 where thread_id = ? `, reqBody.Thread_id)
		// row 추가
		return nil
	}
	// Is_recommended -> is_recommended : false
	if is_recommended == true {
		_, err = db.Exec(`Update thread_recommend set is_recommended = false where thread_id = ? and email = ?`, reqBody.Thread_id, reqBody.Email)
		if err := ErrChecker.Check(err); err != nil {
			return err
		}
		_, err = db.Exec(`Update thread set like_count = like_count - 1 where thread_id = ? `, reqBody.Thread_id)
		return nil
	}
	// Not is_recommneded -> is_recommend : true
	_, err = db.Exec(`Update thread_recommend set is_recommended = true where thread_id = ? and email = ?`, reqBody.Thread_id, reqBody.Email)
	if err := ErrChecker.Check(err); err != nil {
		return err
	}
	_, err = db.Exec(`Update thread set like_count = like_count + 1 where thread_id = ? `, reqBody.Thread_id)
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
