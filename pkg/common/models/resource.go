package models

import "gorm.io/gorm"
/Here is a golang Struct
type Resource struct {
	gorm.Model
	Name        string `json:"title"`
	Description string `json:"description"`
}
