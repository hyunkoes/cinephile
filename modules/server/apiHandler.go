package server

import (
	. "cinephile/modules/api"

	"github.com/gin-gonic/gin"
)

// Thread CRUD
func getThreads(c *gin.Context) {
	threads, err := GetThreadsWithRecommend(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"error": nil, "threads": threads})
	}
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
