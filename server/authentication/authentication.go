package authentication

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginInfo struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type SignupInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

var users = []SignupInfo{
	{Username: "Bill", Password: "123", Email: "Bill@yahoo.com"},
}

func AuthenticationRequests(router *gin.Engine) {
	router.GET("/users", getAllUsers)
	router.GET("/users/:name", getUser)
	router.POST("/users", postUser)
}

func postUser(c *gin.Context) {
	var newUser SignupInfo

	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "unabled to create user", "success": false})
		return
	}
	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func getAllUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func getUser(c *gin.Context) {
	username := c.Param("name")

	for _, u := range users {
		if u.Username == username {
			c.IndentedJSON(http.StatusOK, u)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "unable to find user", "success": false})
}
