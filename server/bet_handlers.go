package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *ApiServer) HandleGetBets(c *gin.Context) {

	bets, err := s.store.GetBets()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bets)

}
