package controller

import (
	"github.com/gin-gonic/gin"
	"moll-y.io/doro/internal/doro/rest/dto"
	"moll-y.io/doro/internal/pkg/service"
	"net/http"
)

type UserController struct {
	Router      gin.IRoutes
	UserService *service.UserService
}

func (uc *UserController) Route() {
	uc.Router.POST("/users", uc.CreateUser)
	uc.Router.GET("/users", uc.FindUser)
}

func (uc *UserController) CreateUser(c *gin.Context) {
	body := dto.CreateUserRequestDto{}
	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	user, err := uc.UserService.CreateUser(body.Name, body.Email, body.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (uc *UserController) FindUser(c *gin.Context) {
	email := c.Query("email")
	user, err := uc.UserService.FindUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
