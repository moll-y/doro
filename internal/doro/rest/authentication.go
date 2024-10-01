package controller

import (
	"github.com/gin-gonic/gin"
	"moll-y.io/doro/internal/doro/rest/dto"
	"moll-y.io/doro/internal/pkg/service"
	"net/http"
)

type AuthenticationController struct {
	Router                gin.IRoutes
	AuthenticationService *service.AuthenticationService
}

func (ac *AuthenticationController) Route() {
	ac.Router.POST("/authentications", ac.Authenticate)
}

func (ac *AuthenticationController) Authenticate(c *gin.Context) {
	request := dto.AuthenticationRequestDto{}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request."})
		return
	}
	t, err := ac.AuthenticationService.Authenticate(request.Email, request.Password)
	if err != nil {
		c.JSON(200, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"jwt": t})
}
