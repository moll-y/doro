package controller

import (
	"github.com/gin-gonic/gin"
	"moll-y.io/doro/internal/pkg/service"
	"net/http"
)

type BacklogController struct {
	Router         gin.IRoutes
	BacklogService *service.BacklogService
}

func (bc *BacklogController) Route() {
	bc.Router.GET("/backlogs", bc.FindBacklogs)
}

func (bc *BacklogController) FindBacklogs(c *gin.Context) {
	backlogs, err := bc.BacklogService.FindBacklogs()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, backlogs)
}
