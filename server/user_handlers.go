package server

import (
	"api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *ApiServer) HandleGetUsers(c *gin.Context) {
	users, err := s.store.GetUsers()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}

func (s *ApiServer) HandleGetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id needs to be an int"})
	}

	user, err := s.store.GetUserById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func (s *ApiServer) HandleCreateUser(c *gin.Context) {

	var newUserReq models.CreateUserRequest

	if err := c.BindJSON(&newUserReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request. You need to pass a user_name, email and a password only"})
		return
	}

	newUser := models.NewUser(newUserReq.UserName, newUserReq.Email, newUserReq.Password)

	if err := s.store.CreateUser(newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newUser)

}

func (s *ApiServer) HandleDeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id needs to be an int"})
	}

	if err := s.store.DeleteUser(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"approved": "The user got deleted correctly"})

}
