package models

import "time"

type Category struct {
	ID        int               `json:"id" gorm:"primary_key:auto_increment"`
	Name      string            `json:"name"`
	Product   []ProductResponse `json:"product"`
	CreatedAt time.Time         `json:"-"`
	UpdatedAt time.Time         `json:"-"`
}

type CategoryResponse struct {
	ID        int               `json:"id"`
	Name      string            `json:"name"`
	ProductID int               `json:"id_product"`
	Product   []ProductResponse `json:"product"`
}

func (CategoryResponse) TableName() string {
	return "category"
}
