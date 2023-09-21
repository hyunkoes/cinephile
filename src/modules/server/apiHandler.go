package server

import (
	. "cinephile/modules/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get thread list
// @Description Get thread list with two parameter. Cursor is pagination id. Type is sorting filter.
// @Accept json
// @Produce json
// @Success 200 {array} Thread
//
//	@Param        cursor    query     string  false  "Pagination id"  Format(number)
//	@Param        type    query     string  false  "Sorting filter" Format(new || hot)
//
// @Router /list/threads [get]
func getThreads(c *gin.Context) {
	_, valid := c.GetQuery(`parent_id`)
	var threads []Thread
	var err error
	var cursor int
	if !valid {
		threads, err, cursor = GetThreadsWithRecommend(c)
	} else {
		threads, err, cursor = GetChildThreadsWithRecommend(c)
	}
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if cursor > 0 {
		c.JSON(200, gin.H{"error": nil, "threads": threads, "count": len(threads), "lastCursor": cursor})
		return
	}
	c.JSON(200, gin.H{"error": nil, "threads": threads, "count": len(threads), "lastCursor": nil})

}

// @Summary Get single thread info
// @Description Get thread with thread_id.
// @Accept json
// @Produce json
// @Success 200 {object} Thread
// @Param        thread_id    query     string  true  "Thread id"  Format(number)
// @Router /threads [get]
func getThread(c *gin.Context) {
	thread, err := GetThread(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"error": nil, "thread": thread})
	}
}

// @Summary Regist thread
// @Description Regist thread.
// @Description If re-thread, should contain parentId ( Parent thread id ).
// @Description Required access token.
// @Accept json
// @Produce json
// @Success 200 {string} error
// @Param thread body ThreadRegistForm true "Thread data"
// @Param user header string true "User token" default(chs)
// @Router /threads [post]
func registThread(c *gin.Context) {
	err := RegistThread(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"error": nil})
	}
}

// @Summary Change like-state of thread
// @Description First like -> create history
// @Description If you cancel like-state, change row state ( Soft delete )
// @Accept json
// @Produce json
// @Success 200 {string} error
// @Param thread body RecommendForm true "Recommend data"
// @Router /threads/like [post]
func changeRecommendThread(c *gin.Context) {
	err := ChangeRecommendThread(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"error": nil})
	}
}

// @Summary Get single movie info
// @Accept json
// @Produce json
// @Success 200 {object} Movie
// @Param movie_id query string true "Movie id"
// @Router /movies [get]
func getMovie(c *gin.Context) {
	movie, err := GetMovie(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"error": nil, "movie": movie})
	}
}

// @Summary Search movie by title ( regexp )
// @Accept json
// @Produce json
// @Success 200 {array} MovieSearch
// @Param keyword query string true "Title keyword"
// @Router /movies/search [get]
func searchMovie(c *gin.Context) {
	movies, cursor, err := SearchMovie(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if cursor > 0 {
		c.JSON(200, gin.H{"error": nil, "movies": movies, "count": len(movies), "lastCursor": cursor})
		return
	}
	c.JSON(200, gin.H{"error": nil, "movies": movies, "count": len(movies), "lastCursor": nil})

}

// test API
func getMovies(c *gin.Context) {
	movie, err := GetMovies(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"error": nil, "movies": movie})
	}
}

// @Summary Get channel list
// @Accept json
// @Produce json
// @Success 200 {array} Channel
//
//	@Param        cursor    query     string  false  "Pagination id"  Format(number)
//	@Param        type    query     string  false  "Sorting filter" Format(new || hot)
//
// @Router /list/channel [get]
func getChannels(c *gin.Context) {
	channels, err := GetChannels(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"error": nil, "channels": channels})
	}
}

// @Summary Get single channel info
// @Description Get thread with channel_id
// @Accept json
// @Produce json
// @Success 200 {object} Channel
// @Param        channel_id    query     string  true  "Channel id"  Format(number)
// @Router /channel [get]
func getChannel(c *gin.Context) {
	channel, err := GetChannel(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"error": nil, "channel": channel})
	}
}

// @Summary Delete thread
// @Description Required authentication of thread.
// @Accept json
// @Produce json
// @Success 200 {object} Thread
// @Param        thread_id    query     string  true  "Thread id"  Format(number)
// @Router /threads [delete]
func deleteThread(c *gin.Context) {
	err := DeleteThread(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"error": nil})
	}
}

func registUser(c *gin.Context) {
	err := RegistUserForTest(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"error": nil})
	}
}

// @Summary Get Hot 8 movies
// @Description
// @Accept json
// @Produce json
// @Success 200 {object} Movie
// @Router /movies/hot [get]
func getHotMovies(c *gin.Context) {
	movies, err := GetHotMovies(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"error": nil, "movies": movies})
	}
}

// @Summary OAuth Login
// @Description Kakao, Google oauth login
// @Accept json
// @Produce json
// @Success 200 {object} Token
// @Router /oauth/callback [get]
func oAuthLogin(c *gin.Context) {
	tokens, err := OAuthLogin(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		referer_domain, valid := c.GetQuery(`state`)
		if !valid {
			referer_domain = "cinephile.site"
		}
		c.SetCookie("access_token", tokens.AccessToken, tokens.Expire, "/", ".cinephile.site", false, true)
		c.SetCookie("refresh_token", tokens.RefreshToken, tokens.RefreshExpire, "/", ".cinephile.site", false, true)
		c.Redirect(http.StatusFound, referer_domain)
	}
}