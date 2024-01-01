package controller

import (
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("views/*.html")
	r.GET("/list", ShowAllTask)
	r.GET("/task/:id", ShowTask)
	r.GET("/task/new", NewTaskForm)
	r.POST("/task/new", CreateTask)
	r.GET("/task/edit/:id", EdittaskForm)
	r.POST("/task/edit/:id", UpdateTask)
	r.GET("/task/delete/:id", DeleteTask)
	return r
}