package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" gorm:"varchar(255)"`
	Email     string    `json:"email" gorm:"varchar(255)"`
	Password  string    `json:"password" gorm:"varchar(255)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
