package handler

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"

	"github.com/Dimoonevs/task-api/internal/model"
	"github.com/Dimoonevs/task-api/internal/service"
	"github.com/gin-gonic/gin"
)

type TaskHandler struct{ svc *service.TaskService }

func NewTaskHandler(s *service.TaskService) *TaskHandler { return &TaskHandler{svc: s} }

func (h *TaskHandler) Register(r *gin.Engine) {
	taskGroup := r.Group("/tasks")
	{
		taskGroup.POST("", h.createTask)
		taskGroup.GET("", h.listTasks)
		taskGroup.GET("/:id", h.getTask)
		taskGroup.PUT("/:id", h.updateTask)
		taskGroup.DELETE("/:id", h.deleteTask)
	}
}

func (h *TaskHandler) createTask(c *gin.Context) {
	var req model.Task
	if err := c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("createTask: bad JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.svc.Create(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, req)
}

func (h *TaskHandler) getTask(c *gin.Context) {
	id := c.Param("id")
	task, err := h.svc.Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if task == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) listTasks(c *gin.Context) {
	status := c.Query("status")

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))

	tasks, err := h.svc.List(model.Status(status), page, size)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) updateTask(c *gin.Context) {
	id := c.Param("id")

	var req model.Task
	if err := c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("updateTask: invalid JSON id=%s", id)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.svc.Update(id, &req); err != nil {
		if err.Error() == "task not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, req)
}

func (h *TaskHandler) deleteTask(c *gin.Context) {
	id := c.Param("id")
	if err := h.svc.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id, "deleted": true})
}
