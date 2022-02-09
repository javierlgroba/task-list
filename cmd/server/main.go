package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/javierlgroba/task-list/pkg/rtTaskStorage"
	"github.com/javierlgroba/task-list/pkg/task"
)

var tasks rtTaskStorage.RtTaskStorage

func getTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks.GetAll())
}

func getTask(c *gin.Context) {
	task_id := c.Param("task_id")
	if val, ok := tasks.GetTask(task_id); ok {
		c.IndentedJSON(http.StatusOK, val)
	} else {
		c.Status(http.StatusNotFound)
	}
}

func removeTasks(c *gin.Context) {
	value := http.StatusNotFound
	if ok := tasks.RemoveAll(); ok {
		value = http.StatusOK
	}
	c.Status(value)
}

func removeTask(c *gin.Context) {
	task_id := c.Param("task_id")
	http_status := http.StatusNotFound
	if ok := tasks.Remove(task_id); ok {
		http_status = http.StatusOK
	}
	c.Status(http_status)
}

func addTask(c *gin.Context) {
	task_text := c.Query("text")
	task_id := uuid.New().String()
	task := task.Task{ID: task_id, Text: task_text}
	if ok := tasks.Add(task); ok {
		c.IndentedJSON(http.StatusCreated, task)
	}
	c.Status(http.StatusNotFound)
}

func main() {
	tasks = rtTaskStorage.NewRtTaskStorage()
	router := gin.Default()

	router.GET("/task/get", getTasks)
	router.GET("/task/get/:task_id", getTask)

	router.POST("/task/remove", removeTasks)
	router.POST("/task/remove/:task_id", removeTask)

	router.POST("/task/add", addTask)

	router.Run("localhost:8080")
}