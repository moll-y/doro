package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"moll-y.io/doro/internal/api/controller"
	"moll-y.io/doro/internal/api/service"
	"moll-y.io/doro/internal/infrastructure/repository"
)

func main() {
	db, err := gorm.Open(sqlite.Open("doro.db"), &gorm.Config{})
	if err != nil {
		panic("failde to connect database")
	}
	r := gin.Default()
	ur := &repository.UserRepository{DB: db}
	us := &service.UserService{UserRepository: ur}
	uc := &controller.UserController{Router: r, UserService: us}
	as := &service.AuthenticationService{UserRepository: ur}
	ac := &controller.AuthenticationController{Router: r, AuthenticationService: as}
	uc.Route()
	ac.Route()
	r.Run()
}
