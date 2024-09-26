package controller

import (
	"github.com/gin-gonic/gin"
	"moll-y.io/doro/internal/api/service"
)

type UserController struct {
	Router      *gin.Engine
	UserService *service.UserService
}

func (uc *UserController) Route() {
	uc.Router.POST("/users", uc.FindUserByEmail)
}

func (uc *UserController) FindUserByEmail(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}
