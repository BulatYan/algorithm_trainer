package handlers

import (
	"context"
	"log"
	"net/http"
	"ped_poject/database"
	"ped_poject/models"
	"time"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	TaskRepo *database.TaskRepository
}

func NewTaskHandler(repo *database.TaskRepository) *TaskHandler {
	return &TaskHandler{TaskRepo: repo}
}

type CreateTaskRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Lvl         int    `json:"lvl" binding:"required"`
}
type SearchTaskRequest struct {
	Name string `json:"name" binding:"required"`
}

type СhangeTaskRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Lvl         int    `json:"lvl" binding:"required"`
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var req CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("invalid request: %v\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	existnameTask, _ := h.TaskRepo.GetTaskByName(context.Background(), req.Name)
	if existnameTask != nil {
		log.Printf("task with this name exists\n")
		c.JSON(http.StatusConflict, gin.H{"error": "task with this name already exists"})
		return
	}
	task := models.Task{
		Name:        req.Name,
		Description: req.Description,
		Lvl:         req.Lvl,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := h.TaskRepo.CreateTask(ctx, &task)
	if err != nil {
		log.Printf("creating task error: %v\n", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "creating task error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
func (h *TaskHandler) SearchTask(c *gin.Context) {
	var req SearchTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("invalid request: %v\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	existname_Task, _ := h.TaskRepo.GetTaskByName(context.Background(), req.Name)
	if existname_Task == nil {
		log.Printf("task not found \n")
		c.JSON(http.StatusConflict, gin.H{"error": "task not found"})
		return
	}
	c.JSON(http.StatusOK, existname_Task)
}
func (h *TaskHandler) Update_Task(c *gin.Context) {
	var req СhangeTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("invalid request: %v\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	task, _ := h.TaskRepo.GetTaskByName(context.Background(), req.Name)
	if task == nil {
		log.Printf("task not found \n")
		c.JSON(http.StatusConflict, gin.H{"error": "task not found"})
		return
	}
	if req.Description != "" {
		task.Description = req.Description
	}
	if req.Lvl != 0 {
		task.Lvl = req.Lvl
	}
	err := h.TaskRepo.UpdateTask(context.Background(), task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "update failed"})
		return
	}

	c.JSON(http.StatusOK, task)
}
