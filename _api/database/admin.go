package database

import (
	"app/models"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func Admin() {
	var user models.User
	DB.First(&user, "username = ?", os.Getenv("admin_username"))

	if user.ID == 0 {
		println("creating default admin user")

		hash, err := bcrypt.GenerateFromPassword([]byte(os.Getenv("admin_password")), 10)

		if err != nil {
			println("failed to create hash")
			return
		}

		user := models.User{Username: os.Getenv("admin_username"), Email: os.Getenv("admin_email"), Password: string(hash), Role: "admin"}
		result := DB.Create(&user)

		if result.Error != nil {
			println("failed to create admin instance")
			return
		}
	}

}
