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

func (s *ApiServer) HandleGetUserBets(c *gin.Context) {
	id, err := getId(c, "user_id")

	if err {
		return
	}

	bets, err2 := s.store.GetUserBets(id)

	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, bets)
}

func (s *ApiServer) HandlePostBetByUser(c *gin.Context) {

	user_id, err := getId(c, "user_id")
	event_id, err2 := getId(c, "event_id")

	if err || err2 {
		return
	}

	var betReq models.CreateBetRequest

	if err3 := c.BindJSON(&betReq); err3 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err3.Error()})
		return
	}

	bet := models.NewBet(user_id, event_id, betReq.Amount, betReq.Choice)

	if err4 := s.store.CreateBet(bet); err4 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err4.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, bet)
}

func (s *ApiServer) HandleDeleteBetByUser(c *gin.Context) {
	user_id, err := getId(c, "user_id")
	event_id, err2 := getId(c, "event_id")

	if err || err2 {
		return
	}

	if err := s.store.DeleteBet(user_id, event_id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bet deleted correctly"})

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
	id, err := getId(c, "user_id")

	if err {
		return
	}

	user, err2 := s.store.GetUserById(id)

	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
		return
	}

	userResponse := &models.UserDTO{
		Id:       user.Id,
		UserName: user.UserName,
		Email:    user.Email,
	}

	c.IndentedJSON(http.StatusOK, userResponse)
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
	id, err := getId(c, "user_id")

	if err {
		return
	}

	if err := s.store.DeleteUserById(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "The user got deleted correctly"})

}

func getId(c *gin.Context, id_string string) (int, bool) {

	id, err := strconv.Atoi(c.Param(id_string))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id needs to be an int"})
		return 0, true
	}

	return id, false
}
