package server

import (
	. "cinephile/modules/api"

	"github.com/gin-gonic/gin"
)

// Thread CRUD
func getThreads(c *gin.Context) {
	posts, err := GetThreads(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"error": nil, "projects": posts})
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
// func getMovie(c *gin.Context) {
// 	movie, err := GetMovie(c)
// 	if err != nil {
// 		c.JSON(400, gin.H{"error": err.Error()})
// 	} else {
// 		c.JSON(200, gin.H{"error": nil, "movie": movie})
// 	}
// }
