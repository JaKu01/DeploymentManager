package db

import (
	"github.com/JaKu01/DeploymentManager/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("db/sqlite/projects.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&model.Project{})
	if err != nil {
		panic("failed to migrate database")
	}
	return db
}
