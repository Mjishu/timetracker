package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Task struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Time     string `json:"time"`
	TaskType string `json:"type"`
	Date     string `json:"date"`
}

var Tasks = []Task{
	Task{Id: "1", Title: "walking the dog", Time: "01:00", TaskType: "exercise", Date: "2024/11/15"},
	Task{Id: "2", Title: "coding this app", Time: "01:30", TaskType: "coding", Date: "2024/11/16"},
}

func handleRequests() {
	router := gin.Default()
	router.GET("/tasks", getTasks)
	router.GET("/tasks/:id", getSpecificTask)
	router.POST("/tasks", postTasks)
	router.DELETE("/tasks/:id", deleteTaskById)

	router.Run("localhost:8080")
}

func getTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Tasks)
}

func postTasks(c *gin.Context) {
	var newTask Task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	Tasks = append(Tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}

func getSpecificTask(c *gin.Context) {
	id := c.Param("id")
	for _, t := range Tasks {
		if t.Id == id {
			c.IndentedJSON(http.StatusOK, t)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}

func deleteTaskById(c *gin.Context) {
	id := c.Param("id")
	for i, t := range Tasks {
		if t.Id == id {
			Tasks = append(Tasks[:i], Tasks[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "task was deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}
