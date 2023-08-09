package repository

import (
	"fmt"
	"gorm.io/gorm"
	//"gorm.io/driver/sqlite"
	service "forummodule/internal/service"
)

// Поиск пользователя по логину и паролю
func TryLogin(db *gorm.DB, user service.AccountData) bool {
	var result service.AccountData
	query := db.Where("login = ? AND password = ?", user.Login, user.Password)

	query.Find(&result)

	// Вывод идентификатора пользователя
	fmt.Println("\n", result)

}

// создание пользователя
func SetUser(db *gorm.DB, user service.AccountData) {

	db.Create(&user)

	db.Save(&user)

	fmt.Println("User created successfully")
}
