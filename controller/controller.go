package controller

import (
	"net/http"

	"task-manager/model"
	"task-manager/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HomePage(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func RegisterUser(c *gin.Context) {
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	err := services.Register(req.Email, req.Password, req.Role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user registered"})
}

func LoginUser(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	token, err := services.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func CreateTask(c *gin.Context) {
	userID := c.MustGet("user_id").(string)

	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	task.UserID, _ = uuid.Parse(userID)
	services.CreateTask(&task)

	c.JSON(http.StatusCreated, task)
}

func GetTasks(c *gin.Context) {
	role := c.MustGet("role").(string)
	userID := c.MustGet("user_id").(string)

	if role == "admin" {
		tasks, _ := services.GetAllTasks()
		c.JSON(http.StatusOK, tasks)
		return
	}

	uid, _ := uuid.Parse(userID)
	tasks, _ := services.GetTasksByUser(uid)
	c.JSON(http.StatusOK, tasks)
}

func DeleteTask(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	services.DeleteTask(id)
	c.JSON(http.StatusOK, gin.H{"message": "task deleted"})
}
