package authentication

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type LoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
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
	router.POST("/users/login", loginUser)
	router.POST("/users/logout", logoutUser)
	router.POST("/users", postUser)
	router.DELETE("/users/:name", deleteUser)
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
		if strings.EqualFold(u.Username, username) {
			c.IndentedJSON(http.StatusOK, u)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "unable to find user", "success": false})
}

func deleteUser(c *gin.Context) {
	username := c.Param("name")

	for i, u := range users {
		if strings.EqualFold(u.Username, username) {
			users = append(users[:i], users[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "user successfully deleted", "success": true})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found", "success": false})
}

func loginUser(c *gin.Context) {
	reqData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "unable to call login user", "success": false})
	}

	var user LoginInfo
	err = json.Unmarshal(reqData, &user)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid json data", "success": false})
	}

	for _, u := range users {
		if strings.EqualFold(u.Username, user.Username) && strings.EqualFold(u.Password, user.Password) {
			fmt.Println("usernames are equal and so are passowords!")
			tokenString, err := createToken(user.Username)
			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error creating token", "success": false})
				return
			}
			// loggedInUser = user.Username
			fmt.Printf("Token created: %s\n", tokenString)
			c.SetCookie("token", tokenString, 3600, "/", "localhost", false, true) //! might not work because this is for only golang, not with go backedn adn js frontend
		} else {
			c.IndentedJSON(http.StatusUnauthorized, "invalid credentials")
		}
	}
}

func logoutUser(c *gin.Context) {
	// more logic for actual logout this is just boilerplate
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	fmt.Println("Logging you out!")
}
