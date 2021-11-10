package rest

import (
	"fmt"
	"net/http"

	"example.com/m/v2/ginrest/user"
	"github.com/gin-gonic/gin"
)

func getUserByID(c *gin.Context) {
	id := c.Param("id")

	for _, user := range user.Users {
		if user.ID == id {
			c.IndentedJSON(http.StatusOK, user)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func AddUserRoutes(router *gin.Engine) {
	router.GET("/users", getUsers)
	router.POST("/users", postUsers)
	router.GET("users/:id", getUserByID)
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, user.Users)
}

var router *gin.Engine

func postUsers(c *gin.Context) {
	var newUser user.User

	if err := c.BindJSON(&newUser); err != nil {
		fmt.Println("Error happened ?")
		// reqBody := c.Request.Response.Body*
		raw, _ := c.GetRawData()
		fmt.Println(string(raw))

		return
	}

	user.Users = append(user.Users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
	// c.IndentedJSON(http.StatusOK,)
}
