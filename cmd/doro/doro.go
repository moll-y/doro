package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"moll-y.io/doro/internal/doro/rest"
	"moll-y.io/doro/internal/doro/rest/middleware"
	"moll-y.io/doro/internal/pkg/domain"
	repository "moll-y.io/doro/internal/pkg/repository/sqlite"
	"moll-y.io/doro/internal/pkg/service"
)

func main() {
	db, err := gorm.Open(sqlite.Open("doro.db"), &gorm.Config{})
	if err != nil {
		panic("failde to connect database")
	}
	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.Organization{})
	ur := &repository.UserRepository{DB: db}
	or := &repository.OrganizationRepository{DB: db}
	us := &service.UserService{UserRepository: ur}
	as := &service.AuthenticationService{UserRepository: ur}
	os := &service.OrganizationService{UserRepository: ur, OrganizationRepository: or}
	r := gin.Default()
	ac := &controller.AuthenticationController{Router: r, AuthenticationService: as}
	oc := &controller.OrganizationController{Router: r, OrganizationService: os}
	uc := &controller.UserController{Router: r, UserService: us}
	am := &middleware.AuthenticationMiddleware{AuthenticationService: as}
	uc.Route()
	ac.Route()
	oc.Route()
	r.Use(am.Middleware())
	r.Run()
}
