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
	body := dto.AuthenticateUserRequestDto{}
	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	t, err := ac.AuthenticationService.Authenticate(body.Email, body.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"jwt": t})
}
