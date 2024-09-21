package model

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name              string `json:"name"`
	Url               string `json:"url"`
	Branch            string `json:"branch"`
	DockerComposePath string `json:"dockerComposePath"`
}
