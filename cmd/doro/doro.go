package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"moll-y.io/doro/internal/doro/rest"
	"moll-y.io/doro/internal/doro/rest/middleware"
	"moll-y.io/doro/internal/pkg/domain"
	repository "moll-y.io/doro/internal/pkg/repository/maria"
	"moll-y.io/doro/internal/pkg/service"
)

func main() {
	dns := "root:root@tcp(127.0.0.1:3306)/doro?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.Task{})
	db.AutoMigrate(&domain.Backlog{})
	db.AutoMigrate(&domain.Pomodoro{})

	ur := &repository.UserRepository{DB: db}
	us := &service.UserService{UserRepository: ur}
	as := &service.AuthenticationService{UserRepository: ur}
	am := &middleware.AuthenticationMiddleware{AuthenticationService: as}
	r := gin.Default()
	r.Use(am.Middleware())
	ac := &controller.AuthenticationController{Router: r, AuthenticationService: as}
	uc := &controller.UserController{Router: r, UserService: us}
	uc.Route()
	ac.Route()
	r.Run()
}
