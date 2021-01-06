package controller

import (
	"bubble/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK,"index.html",nil)
}

func CreateTodo(c *gin.Context) {

	//从请求中把数据拿出来
	var todo models.Todo
	c.BindJSON(&todo)

	//存入数据库
	err := models.CreateATodo(&todo)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{ "error":err.Error() })
	} else {
		c.JSON(http.StatusOK,todo)
	}
}

func GetTodoList(c *gin.Context) {
	todoList,err := models.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusOK,gin.H{ "error":err.Error() })
	} else {
		c.JSON(http.StatusOK,todoList)
	}
}

func UpdateATodo(c *gin.Context) {
	id,ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK,gin.H{ "error":"无效id" })
		return
	}
	todo,err := models.GetATodo(id)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{ "error":err.Error() })
	} else {
		c.JSON(http.StatusOK,todo)
	}
}

func DeleteATodo(c *gin.Context) {
	id,ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK,gin.H{ "error":"无效id" })
		return
	}
	if err := models.DeleteATodo(id);err != nil {
		c.JSON(http.StatusOK,gin.H{ "error":err.Error() })
	} else {
		c.JSON(http.StatusOK,gin.H{id:"deleted" })
	}
}