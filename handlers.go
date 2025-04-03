package main

import (
	"api/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleGetUsers(c *gin.Context) {

}

func HandleCreateUser(c *gin.Context) {

	var userReq types.CreateUserRequest

	if err := c.BindJSON(&userReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request. You need to pass a user_name, email and a password only"})
		return
	}

	newUser := types.NewUser(userReq.UserName, userReq.Email, userReq.Password)

	if err := pStore.CreateUser(newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newUser)

}
