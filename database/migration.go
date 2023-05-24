package database

import (
	"dumbmerch/models"
	"dumbmerch/pkg/mysql"
	"fmt"
)

// automatic migrate
func RunMigration() {
	err := mysql.DB.AutoMigrate(&models.User{})

	if err != nil {
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
