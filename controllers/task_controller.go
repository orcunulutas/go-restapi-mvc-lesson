package controllers

import (
	"go-restapi/database"
	"go-restapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllTasks(c *gin.Context) {
	var tasks []models.Task
	db := database.DB
	db.Find(&tasks)
	c.JSON(http.StatusOK, tasks)

}
