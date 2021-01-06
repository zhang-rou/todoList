package routers

import (
	"bubble/controller"
	"bubble/setting"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	if setting.Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()

	//找静态文件
	r.Static("/static","static")
	//找模版文件
	r.LoadHTMLGlob("templates/*")

	r.GET("/",controller.IndexHandler)

	v1Group := r.Group("v1")
	{
		v1Group.POST("/todo",controller.CreateTodo)
		v1Group.GET("/todo",controller.GetTodoList)
		v1Group.PUT("/todo/:id",controller.UpdateATodo)
		v1Group.DELETE("/todo/:id",controller.DeleteATodo)
	}
	return r

}