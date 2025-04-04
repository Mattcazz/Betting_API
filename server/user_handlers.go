package server

import (
	"api/middleware"
	"api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *ApiServer) HandleLogin(c *gin.Context) {
	var loginReq models.LoginRequest

	if err := c.BindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := s.store.GetUserByEmail(loginReq.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email address"})
		return
	}

	if !user.ValidatePassword(loginReq.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
		return
	}

	token, err2 := middleware.CreateJWTtoken(user)

	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
		return
	}

	loginResponse := &models.LoginResponse{
		UserId: user.Id,
		Token:  token,
	}

	c.IndentedJSON(http.StatusAccepted, loginResponse)

}

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

func (s *ApiServer) HandleDeleteUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id needs to be an int"})
	}

	if err := s.store.DeleteUserById(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"approved": "The user got deleted correctly"})

}
