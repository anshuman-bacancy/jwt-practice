package services

import (
	"fmt"
	"jwt/database"
	"jwt/models"
	"jwt/utils"
	"log"
)

// SaveUser saves to db
func SaveUser(name, email, password, role string) {
	db := database.OpenDb()
	defer database.Closedb(db)
	if db == nil {
		log.Fatal("Db is nil")
	}

	hashedPassword, _ := utils.HashPassword(password)

	user := models.User{Name: name, Email: email, Password: hashedPassword, Role: role}
	db.Select("Name", "Email", "Password", "Role").Create(&user)
}

// SignInUser checks for valid user from db
func SignInUser(email, password string) (bool, models.User) {
	db := database.OpenDb()

	var user models.User
	db.Where("Email", email).First(&user)
	fmt.Println("user is :=> ", user.Email, user.Password)
	match := utils.CheckPasswordHash(password, user.Password)
	return match, user
}
