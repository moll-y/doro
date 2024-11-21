package controller

import (
	"github.com/gin-gonic/gin"
	"moll-y.io/doro/internal/pkg/service"
	"net/http"
)

type TaskController struct {
	Router      gin.IRoutes
	TaskService *service.TaskService
}

func (tc *TaskController) Route() {
	tc.Router.GET("/tasks", tc.FindTasks)
	tc.Router.PUT("/tasks", tc.UpdateTask)
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
	task, err := tc.TaskService.UpdateTask()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (tc *TaskController) FindTasks(c *gin.Context) {
	tasks, err := tc.TaskService.FindTasks()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}
