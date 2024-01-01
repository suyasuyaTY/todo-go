package controller

import (
	"fmt"
	"my-todilist/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowAllTask(ctx *gin.Context) {
	tasks, err := model.Index()
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.HTML(http.StatusOK, "task_list.html", gin.H{"Title": "Task list", "Tasks": tasks})
}

func ShowTask(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	task, err := model.Show(id)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.HTML(http.StatusOK, "task.html", gin.H{"Title": "Task", "Task": task})
}

func NewTaskForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "new_task_form.html", gin.H{"Title": "Task registration"})
}

func CreateTask(ctx *gin.Context) {
	title, exist := ctx.GetPostForm("title")
	if !exist {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "No title is given"})
		return
	}
	description, exist := ctx.GetPostForm("description")
	if !exist {
		description = ""
	}
	taskID, err := model.Create(title, description)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if taskID == -1 {
		ctx.Redirect(http.StatusFound, "/list")
		return
	}
	ctx.Redirect(http.StatusFound, fmt.Sprintf("/task/%d", taskID))
}

func EdittaskForm(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	task, err := model.Show(id)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.HTML(http.StatusOK, "edit_task_form.html", gin.H{"Title": fmt.Sprintf("Edit task %d", task.ID), "Task": task})
}

func UpdateTask(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	title, exist := ctx.GetPostForm("title")
	if !exist {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	isdone, exist := ctx.GetPostForm("is_done")
	if !exist {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	isDone, err := strconv.ParseBool(isdone)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	description, exist := ctx.GetPostForm("description")
	if !exist {
		description = ""
	}
	path, err := model.Update(id, title, isDone, description)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Redirect(http.StatusFound, path)	
}

func DeleteTask(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = model.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Redirect(http.StatusFound, "/list")	
}