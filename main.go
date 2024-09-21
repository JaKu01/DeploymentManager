package main

import (
	"github.com/JaKu01/DeploymentManager/db"
	"github.com/JaKu01/DeploymentManager/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	database := db.InitDatabase()

	r := gin.Default()
	handlers.Setup(r, database)
	r.Run(":8080")
}
