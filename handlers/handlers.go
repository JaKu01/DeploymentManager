package handlers

import (
	"DeploymentManager/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func getProjects(database *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var projects []model.Project
		database.Find(&projects)
		c.JSON(http.StatusOK, projects)
	}
}

func postProjects(database *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var project model.Project
		err := c.BindJSON(&project)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		database.Create(&project)
		c.Status(http.StatusCreated)
	}
}

func putProjects(database *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var project model.Project
		err := c.BindJSON(&project)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var projectFromDb model.Project
		result := database.First(&projectFromDb, project.ID)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
			return
		}

		database.Save(&project)
		c.Status(http.StatusOK)
	}
}

func deleteProjects(database *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var project model.Project
		err := c.BindJSON(&project)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var projectFromDb model.Project
		result := database.First(&projectFromDb, project.ID)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
			return
		}

		database.Delete(&project)
		c.Status(http.StatusOK)
	}
}

func Setup(r *gin.Engine, db *gorm.DB) {
	r.GET("/projects", getProjects(db))
	r.POST("/projects", postProjects(db))
	r.PUT("/projects", putProjects(db))
	r.DELETE("/projects", deleteProjects(db))
}
