package server

import (
	"fmt"
	"os"

	docs "cinephile/docs"
	. "cinephile/modules/storage"

	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerFiles "github.com/swaggo/files"

	"github.com/gin-gonic/gin"
)

const port = ":4000"

func Serve(mode int) { // local : 4000 호스팅 시작
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api"
	api := r.Group("/api")

	os.Setenv("TZ", "Asia/Seoul")
	if GetConn().Ping() != nil {
		panic(fmt.Errorf("mysql is off status"))
	}

	RegistApiHandler(api)
	r.Run(port)
}

func RegistApiHandler(api *gin.RouterGroup) {
	RegistChannelApiHandler(api)
	RegistMovieApiHandler(api)
	RegistUserApiHandler(api)
	RegistAccountApiHandler(api)
	RegistThreadApiHandler(api)
	RegistSwaggerApiHandler(api)
}
func RegistSwaggerApiHandler(api *gin.RouterGroup) {
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
	// kakao, google 둘 다 여기로 callback, platform -> parameter
	api.GET("/oauth/callback", oAuthLogin)

}
func RegistUserApiHandler(api *gin.RouterGroup) {
	/*  Reply			200 -> thread list
	400 -> No more thread
	*/
	// api.GET("/user", getUser)
	/*  Reply			200 -> threads
	400 -> No more thread
	*/
	api.GET("/users", getUser)
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
	api.GET("/movies", getMovie)

	api.GET("/movies/search", searchMovie)

	api.GET("/movies/hot", getHotMovies)
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
