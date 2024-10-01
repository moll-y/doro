package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"moll-y.io/doro/internal/doro/rest"
	"moll-y.io/doro/internal/doro/rest/middleware"
	repository "moll-y.io/doro/internal/pkg/repository/sqlite"
	"moll-y.io/doro/internal/pkg/service"
)

func main() {
	db, err := gorm.Open(sqlite.Open("doro.db"), &gorm.Config{})
	if err != nil {
		panic("failde to connect database")
	}
	router := gin.Default()
	router.Use(middleware.AuthenticationMiddleware())
	ur := &repository.UserRepository{DB: db}
	us := &service.UserService{UserRepository: ur}
	uc := &controller.UserController{Router: router, UserService: us}
	as := &service.AuthenticationService{UserRepository: ur}
	ac := &controller.AuthenticationController{Router: router, AuthenticationService: as}
	uc.Route()
	ac.Route()
	router.Run()
}
