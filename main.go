package main

import (
	"DeploymentManager/db"
	"DeploymentManager/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	database := db.InitDatabase()

	r := gin.Default()
	handlers.Setup(r, database)
	r.Run(":8080")
}
