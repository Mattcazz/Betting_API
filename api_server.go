package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ApiServer struct {
	engine *gin.Engine
}

func NewApiServer() *ApiServer {
	r := gin.Default()

	return &ApiServer{
		engine: r,
	}
}

func (s *ApiServer) SetUpRoutes() {
	s.SetUpUserRoutes()
	fmt.Println("Routes are set up correctly")
}

func (s *ApiServer) SetUpUserRoutes() {
	userGroup := s.engine.Group("/user")
	{
		userGroup.GET("", HandleGetUsers)
		userGroup.GET("/:id", HandleGetUserById)
		userGroup.POST("", HandleCreateUser)

	}
}
