package server

import (
	"api/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *ApiServer) HandleGetBet(c *gin.Context) {
	user_id, event_id, err := fetchUserIdEventId("user_id", "event_id", c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bet, err := s.store.GetBet(user_id, event_id)

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

	user_id, event_id, err := fetchUserIdEventId("user_id", "event_id", c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bet := models.NewBet(user_id, event_id, newBetReq.Amount, newBetReq.Choice)

	if err := s.store.CreateBet(bet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

func (s *ApiServer) HandleDeleteBet(c *gin.Context) {
	user_id, event_id, err := fetchUserIdEventId("user_id", "event_id", c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := s.store.DeleteBet(user_id, event_id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Bet deleted correctly"})
}

func fetchUserIdEventId(user_id, event_id string, c *gin.Context) (int, int, error) {
	userId, err := strconv.Atoi(c.Param(user_id))

	if err != nil {
		return 0, 0, fmt.Errorf("user_id not valid. It has to be an int")
	}

	eventId, err2 := strconv.Atoi(c.Param(event_id))

	if err2 != nil {
		return 0, 0, fmt.Errorf("user_id not valid. It has to be an int")
	}

	return userId, eventId, nil
}
