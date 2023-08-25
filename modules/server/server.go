package server

import (
	"fmt"

	. "cinephile/modules/storage"

	"github.com/gin-gonic/gin"
)

const port = ":4000"

func Serve(mode int) { // local : 4000 호스팅 시작
	r := gin.Default()
	if GetConn().Ping() != nil {
		panic(fmt.Errorf("mysql is off status"))
	}
	api := r.Group("/api")
	RegistApiHandler(api)
	r.Run(port)
}

func RegistApiHandler(api *gin.RouterGroup) {
	RegistChannelApiHandler(api)
	RegistMovieApiHandler(api)
	RegistUserApiHandler(api)
	RegistThreadApiHandler(api)
}
func RegistThreadApiHandler(api *gin.RouterGroup) {
	/*  Reply			200 -> thread list
	400 -> No more thread
	*/
	api.GET("/list/threads", getThreads)
	/*  Reply			200 -> thread detail
	400 -> No more thread
	*/
	api.GET("/threads", getThread)
	/*  Reply			200 -> threads
	400 -> No more thread
	*/
	api.POST("/threads", registThread)
	/*  Reply			200 -> threads
	400 -> No more thread
	*/
	api.POST("/threads/likes", changeRecommendThread)
	/*  Reply			200 -> threads
	400 -> No more thread
	*/
	api.DELETE("/threads", deleteThread)
	/*  Reply			200 -> like thread
	else -> unknown error
	*/
	// api.("/thread", likeThread)

}
func RegistAccountApiHandler(api *gin.RouterGroup) {
	/*  Reply			200 -> thread list
	400 -> No more thread
	*/
	// api.POST("/signup", signUp)
	/*  Reply			200 -> thread list
	400 -> No more thread
	*/
	// api.POST("/signin", signIn)
	/*  Reply			200 -> thread list
	400 -> No more thread
	*/
	// api.PUT("/reset", resetAccount)
	/*  Reply			200 -> thread list
	400 -> No more thread
	*/
	// api.POST("/auth", autheticate)

}
func RegistUserApiHandler(api *gin.RouterGroup) {
	/*  Reply			200 -> thread list
	400 -> No more thread
	*/
	// api.GET("/user", getUser)
	/*  Reply			200 -> threads
	400 -> No more thread
	*/
	// api.PUT("/user", updateUser)
	/*  Reply			200 -> thread list
	400 -> No more thread
	*/
	// api.GET("/user/like/:id", changeLikeState)
	/*  Reply			200 -> thread list
	400 -> No more thread
	*/
	// api.GET("/user/subscribe/:id", changeSubscribeState)

}

func RegistMovieApiHandler(api *gin.RouterGroup) {
	/*  Reply			200 -> thread list
	400 -> No more thread
	*/
	// api.GET("/movie", getMovie)
	api.GET("/movies/search", searchMovie)
}

func RegistChannelApiHandler(api *gin.RouterGroup) {
	/*  Reply			200 -> thread list
	400 -> No more thread
	*/
	api.GET("/channels", getChannel)
	/*  Reply			200 -> thread detail
	400 -> No more thread
	*/
	api.GET("/list/channels", getChannels)
	/*  Reply			200 -> threads
	400 -> No more thread
	*/
	// api.POST("/channel", registThread)
}
