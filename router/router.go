package router

import (
	"github.com/baaj2109/webcam_server/api"
	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {
	// group1 := engine.Group("/web")
	// {
	// 	group1.GET("/home", api.GetHome)
	// }
	engine.GET("/home", api.GetHome)

	engine.GET("/start_webcam", api.StartWebCam)
	engine.GET("/list", api.ListAllCamera)
	engine.GET("/stop_webcam", api.StopWebCam)
	engine.POST("/set_cam/:webcam", api.SelectCamera)

	authGroup := engine.Group("/auth")
	{
		authGroup.POST("/login", api.Login)
		authGroup.GET("/logout", api.Logout)
		authGroup.POST("/register", api.Register)
	}

}
