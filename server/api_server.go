package server

import (
	"api/store"
	"fmt"

	"github.com/gin-gonic/gin"
)

type ApiServer struct {
	Engine *gin.Engine
	store  store.Store
}

func NewApiServer(store store.Store) *ApiServer {
	r := gin.Default()

	return &ApiServer{
		Engine: r,
		store:  store,
	}
}

func (s *ApiServer) SetUpRoutes() {
	s.SetUpUserRoutes()
	s.SetupEventRoutes()
	s.SetupBetRoutes()
	fmt.Println("Routes are set up correctly")
}

func (s *ApiServer) SetUpUserRoutes() {
	userGroup := s.Engine.Group("/users")
	{
		userGroup.GET("", s.HandleGetUsers)
		userGroup.GET("/:id", s.HandleGetUserById)
		userGroup.POST("", s.HandleCreateUser)
		userGroup.DELETE("/:id", s.HandleDeleteUserById)
	}
}

func (s *ApiServer) SetupEventRoutes() {
	eventGroup := s.Engine.Group("/events")
	{
		eventGroup.GET("", s.HandleGetEvents)
		eventGroup.GET("/:id", s.HandleGetEventById)
		eventGroup.POST("", s.HandleCreateEvent)
		eventGroup.DELETE("/:id", s.HandleDeleteEventById)

	}
}

func (s *ApiServer) SetupBetRoutes() {
	betGroup := s.Engine.Group("/bets")
	{
		betGroup.GET("", s.HandleGetBets)
		betGroup.DELETE("/users/:user_id/events/:event_id", s.HandleDeleteBet)
		betGroup.POST("/users/:user_id/events/:event_id", s.HandleCreateBet)
		betGroup.GET("/users/:user_id/events/:event_id", s.HandleGetBet)
	}
}
