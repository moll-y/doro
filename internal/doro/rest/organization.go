package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"moll-y.io/doro/internal/doro/rest/dto"
	"moll-y.io/doro/internal/pkg/service"
	"net/http"
)

type OrganizationController struct {
	Router              gin.IRoutes
	OrganizationService *service.OrganizationService
}

func (oc *OrganizationController) Route() {
	oc.Router.POST("/organizations", oc.CreateOrganization)
}

func (oc *OrganizationController) CreateOrganization(c *gin.Context) {
	actor, ok := c.Get("actor")
	if !ok {
		c.JSON(http.StatusForbidden, gin.H{"message": "Permission denied."})
		return
	}
	cord := dto.CreateOrganizationRequestDto{}
	if err := c.ShouldBindBodyWith(&cord, binding.JSON); err == nil {
		org, err := oc.OrganizationService.CreateOrganization(actor, cord.Name, cord.Description)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, org)
		return
	}
	iord := dto.ImportOrganizationRequestDto{}
	if err := c.ShouldBindBodyWith(&iord, binding.JSON); err == nil {
		org, err := oc.OrganizationService.ImportOrganization(actor, iord.Source)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, org)
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request."})
}
