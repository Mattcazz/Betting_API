package server

import (
	"api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *ApiServer) HandleGetEventById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id needs to be an int"})
	}

	event, err := s.store.GetEventById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, event)

}

func (s *ApiServer) HandleCreateEvent(c *gin.Context) {
	var newEventReq models.CreateEventRequest

	if err := c.BindJSON(&newEventReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newEvent := models.NewEvent(newEventReq.Name, newEventReq.StartTime, newEventReq.Status)

	if err := s.store.CreateEvent(newEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newEvent)
}

func (s *ApiServer) HandleGetEvents(c *gin.Context) {
	events, err := s.store.GetEvents()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, events)
}

func (s *ApiServer) HandleDeleteEventById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id needs to be an int"})
	}

	if err := s.store.DeleteEventById(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"approved": "The event got deleted correctly"})
}
