package handlers

import (
	"github.com/JaKu01/DeploymentManager/model"
	"github.com/JaKu01/DeploymentManager/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func getProjectsHandler(database *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var projects []model.Project
		database.Find(&projects)
		c.JSON(http.StatusOK, projects)
	}
}

func postProjectsHandler(database *gorm.DB) gin.HandlerFunc {
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

func putProjectsHandler(database *gorm.DB) gin.HandlerFunc {
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

		database.Updates(&project)
		c.Status(http.StatusOK)
	}
}

func deleteProjectsHandler(database *gorm.DB) gin.HandlerFunc {
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

func deployProjectHandler(database *gorm.DB) gin.HandlerFunc {
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

		if err := tools.FullDeployment(&projectFromDb); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusOK)
	}
}

func getUiHandler(database *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var projects []model.Project
		database.Find(&projects)
		c.HTML(http.StatusOK, "index.html", projects)
	}
}

func Setup(r *gin.Engine, db *gorm.DB) {
	r.LoadHTMLGlob("templates/*.html")

	r.GET("/projects", getProjectsHandler(db))
	r.POST("/projects", postProjectsHandler(db))
	r.PUT("/projects", putProjectsHandler(db))
	r.DELETE("/projects", deleteProjectsHandler(db))

	r.POST("/deploy", deployProjectHandler(db))
	r.GET("/", getUiHandler(db))
	r.StaticFile("/create", "./templates/create.html")
	r.StaticFile("/edit", "./templates/edit.html")

	r.StaticFile("/openapi.yaml", "./docs/openapi.yaml")
	r.StaticFile("/swagger", "./docs/index.html")
}
