package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mjishu/timetracker/authentication"
	"github.com/mjishu/timetracker/tasks"
)

func handleRequests() {
	router := gin.Default()

	tasks.HandleTaskRequests(router)
	authentication.AuthenticationRequests(router)

	router.Run("localhost:8080")
}
