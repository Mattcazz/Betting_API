package server

import (
	"api/middleware"
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
	s.Engine.POST("/login", s.HandleLogin)
	s.SetUpUserRoutes()
	s.SetupEventRoutes()
	s.SetupBetRoutes()
	fmt.Println("Routes are set up correctly")
}

func (s *ApiServer) SetUpUserRoutes() {
	userGroup := s.Engine.Group("/users")
	{
		userGroup.GET("", s.HandleGetUsers)
		userGroup.GET("/:user_id", middleware.JWTAuth(s.HandleGetUserById, s.store))
		userGroup.GET("/:user_id/bets", middleware.JWTAuth(s.HandleGetUserBets, s.store))
		userGroup.POST("", s.HandleCreateUser)
		userGroup.POST("/:user_id/events/:event_id/bet", middleware.JWTAuth(s.HandlePostBetByUser, s.store))
		userGroup.DELETE("/:user_id/events/:event_id/bet", middleware.JWTAuth(s.HandleDeleteBetByUser, s.store))
		userGroup.DELETE("/:user_id", middleware.JWTAuth(s.HandleDeleteUserById, s.store))
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
	}
}
