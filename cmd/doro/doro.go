package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"moll-y.io/doro/internal/doro/rest"
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

	users := []domain.User{
		{Name: "Alice", Email: "alice@example.com", Password: "securepassword1"},
		{Name: "Bob", Email: "bob@example.com", Password: "securepassword2"},
		{Name: "Charlie", Email: "charlie@example.com", Password: "securepassword3"},
	}
	db.Create(&users)

	backlogs := []domain.Backlog{
		{NumberOfTasksAllowedToCreate: 100},
		{NumberOfTasksAllowedToCreate: 100},
		{NumberOfTasksAllowedToCreate: 100},
	}
	db.Create(&backlogs)

	tasks := []domain.Task{
		{Name: "...", State: "Pending"},
		{Name: "...", State: "Pending"},
		{Name: "...", State: "Pending"},
	}
	db.Create(&tasks)

	r := gin.Default()

	br := &repository.BacklogRepository{DB: db}
	bs := &service.BacklogService{BacklogRepository: br}
	bc := &controller.BacklogController{Router: r, BacklogService: bs}
	bc.Route()

	tr := &repository.TaskRepository{DB: db}
	ts := &service.TaskService{TaskRepository: tr}
	tc := &controller.TaskController{Router: r, TaskService: ts}
	tc.Route()

	ur := &repository.UserRepository{DB: db}
	us := &service.UserService{UserRepository: ur}
	uc := &controller.UserController{Router: r, UserService: us}
	uc.Route()

	//as := &service.AuthenticationService{UserRepository: ur}
	//am := &middleware.AuthenticationMiddleware{AuthenticationService: as}
	//ac := &controller.AuthenticationController{Router: r, AuthenticationService: as}
	//r.Use(am.Middleware())

	//ac.Route()
	r.Run()
}
