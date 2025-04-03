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
	fmt.Println("Routes are set up correctly")
}

func (s *ApiServer) SetUpUserRoutes() {
	userGroup := s.Engine.Group("/user")
	{
		userGroup.GET("", s.HandleGetUsers)
		userGroup.GET("/:id", s.HandleGetUserById)
		userGroup.POST("", s.HandleCreateUser)
		userGroup.DELETE("/:id", s.HandleDeleteUser)
	}
}
