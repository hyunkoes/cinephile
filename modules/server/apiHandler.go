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
