package router

import (
	"github.com/gin-gonic/gin"
	"rankwillServer/controller"
	"rankwillServer/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)
	r.POST("/api/querypage", controller.Getpage)
	r.POST("/api/querybyname", controller.Getbyname)
	r.POST("/api/auth/follow", middleware.AuthMiddleware(), controller.Follow)
	r.POST("/api/auth/unfollow", middleware.AuthMiddleware(), controller.Unfollow)
	r.POST("/api/auth/getfollowing", middleware.AuthMiddleware(), controller.GetFollowing)
	r.POST("/api/getcontestbypage", controller.GetContest)
	r.POST("/api/auth/getfollowlist", middleware.AuthMiddleware(), controller.GetFollowList)
	return r
}
