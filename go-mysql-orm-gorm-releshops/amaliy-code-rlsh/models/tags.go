package models

import "gorm.io/gorm"



type Tags struct {
	gorm.Model
	Name  string  `gorm:"unique;not null" validate:"required"`
	Posts []Post  `gorm:"many2many:post_tags;"`
}