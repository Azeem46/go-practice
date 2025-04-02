package controllers

import (
	"go-practice/config"
	"go-practice/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create Task
func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&task)
	c.JSON(http.StatusOK, task)
}

// Get All Tasks
func GetTasks(c *gin.Context) {
	var tasks []models.Task
	config.DB.Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}

// Update Task
func UpdateTask(c *gin.Context) {
	var task models.Task
	id := c.Param("id")
	if err := config.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.ShouldBindJSON(&task)
	config.DB.Save(&task)
	c.JSON(http.StatusOK, task)
}

// Delete Task
func DeleteTask(c *gin.Context) {
	var task models.Task
	id := c.Param("id")
	if err := config.DB.Delete(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
