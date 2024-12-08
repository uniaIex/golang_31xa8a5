package database

import "app/models"

func Sync() {
	DB.AutoMigrate(&models.User{}, &models.Project{}, &models.File{})
}
