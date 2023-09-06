package server

import (
	. "cinephile/modules/api"

	"github.com/gin-gonic/gin"
)

// Thread CRUD
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
func getThread(c *gin.Context) {
	thread, err := GetThread(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"error": nil, "thread": thread})
	}
}
func registThread(c *gin.Context) {
	err := RegistThread(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"error": nil})
	}
}
func changeRecommendThread(c *gin.Context) {
	err := ChangeRecommendThread(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"error": nil})
	}
}

// Movie R
func getMovie(c *gin.Context) {
	movie, err := GetMovie(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"error": nil, "movie": movie})
	}
}

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

// Movie R
func getMovies(c *gin.Context) {
	movie, err := GetMovies(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"error": nil, "movies": movie})
	}
}

/*
	 Channel list
		@ type = new / hot
			@ new is ordering latest thread
			@ hot is ordering hotest thread
*/
func getChannels(c *gin.Context) {
	channels, err := GetChannels(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"error": nil, "channels": channels})
	}
}
func getChannel(c *gin.Context) {
	channel, err := GetChannel(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"error": nil, "channel": channel})
	}
}

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

func getHotMovies(c *gin.Context) {
	movies, err := GetHotMovies(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"error": nil, "movies": movies})
	}
}
