package main

import (
	"log"
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
  router := gin.Default()
  router.Use(func (c *gin.Context) {
    header := c.GetHeader("Authorization")
    if !strings.HasPrefix(header, "Bearer ") {
      log.Println("no bearer")
      return
    }
    jwt := strings.TrimPrefixeaderauthHeader, "Bearer ")
    c.Set("jwt", jwt)
    c.Next()
  })
	ur := &repository.UserRepository{DB: db}
	us := &service.UserService{UserRepository: ur}
	uc := &controller.UserController{Router: router, UserService: us}
	as := &service.AuthenticationService{UserRepository: ur}
	ac := &controller.AuthenticationController{Router: router, AuthenticationService: as}
	uc.Route()
	ac.Route()
	router.Run()
}
