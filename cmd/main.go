package main

import (
	"api_rest_go/internal/handler"
	"api_rest_go/internal/repository"
	"api_rest_go/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	repo := repository.NewInMemoryUserRepository()
	svc := service.NewUserService(repo)
	hdl := handler.NewUserHandler(svc)

	router := gin.Default()
	hdl.RegisterRoutes(router.Group(""))

	router.Run(":8080")
}
