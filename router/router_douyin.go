package router

import (
	"github.com/gin-gonic/gin"
	"github.com/handlers/comment"
	"github.com/handlers/user_info"
	"github.com/handlers/user_login"
	"github.com/handlers/video"
	"github.com/middleware"
	"github.com/models"
)

func InitDouyinRouter() *gin.Engine {
	models.InitDB()
	r := gin.Default()

	r.Static("static", "./static")

	baseGroup := r.Group("/douyin")
	//根据灵活性考虑是否加入JWT中间件来进行鉴权，还是在之后再做鉴权
	// basic apis
	baseGroup.GET("/feed/", video.FeedVideoListHandler)
	baseGroup.GET("/user/", middleware.JWTMiddleWare(), user_info.UserInfoHandler)
	baseGroup.POST("/user/login/", middleware.SHAMiddleWare(), user_login.UserLoginHandler)
	baseGroup.POST("/user/register/", middleware.SHAMiddleWare(), user_login.UserRegisterHandler)
	baseGroup.POST("/publish/action/", middleware.JWTMiddleWare(), video.PublishVideoHandler)
	baseGroup.GET("/publish/list/", middleware.NoAuthToGetUserId(), video.QueryVideoListHandler)

	//extend 1
	baseGroup.POST("/favorite/action/", middleware.JWTMiddleWare(), video.PostFavorHandler)
	baseGroup.GET("/favorite/list/", middleware.NoAuthToGetUserId(), video.QueryFavorVideoListHandler)
	baseGroup.POST("/comment/action/", middleware.JWTMiddleWare(), comment.PostCommentHandler)
	baseGroup.GET("/comment/list/", middleware.JWTMiddleWare(), comment.QueryCommentListHandler)

	//extend 2
	baseGroup.POST("/relation/action/", middleware.JWTMiddleWare(), user_info.PostFollowActionHandler)
	baseGroup.GET("/relation/follow/list/", middleware.NoAuthToGetUserId(), user_info.QueryFollowListHandler)
	baseGroup.GET("/relation/follower/list/", middleware.NoAuthToGetUserId(), user_info.QueryFollowerHandler)
	return r
}
