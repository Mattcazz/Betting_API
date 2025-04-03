package server

import (
	"api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *ApiServer) HandleGetBetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id not valid. It has to be an int"})
		return
	}

	bet, err := s.store.GetBetById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bet)
}

func (s *ApiServer) HandleCreateBet(c *gin.Context) {
	var newBetReq models.CreateBetRequest

	if err := c.BindJSON(&newBetReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "There was an error parsing the request. Check your JSON"})
		return
	}

	user_id, err := strconv.Atoi(c.Param("user_id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id not valid. It has to be an int"})
		return
	}

	event_id, err2 := strconv.Atoi(c.Param("event_id"))

	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "event_id not valid. It has to be an int"})
		return
	}

	bet := models.NewBet(user_id, event_id, newBetReq.Amount, newBetReq.Choice)

	if err3 := s.store.CreateBet(bet); err3 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err3.Error()})
		return
	}

	c.JSON(http.StatusCreated, bet)

}

func (s *ApiServer) HandleGetBets(c *gin.Context) {

	bets, err := s.store.GetBets()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bets)

}

func (s *ApiServer) HandleDeleteBetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id not valid. It has to be an int"})
		return
	}

	if err := s.store.DeleteBetById(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Bet deleted correctly"})
}
