package models

import "gorm.io/gorm"

type Resource struct {
	gorm.Model
	Name        string `json:"title"`
	Description string `json:"description"`
}
