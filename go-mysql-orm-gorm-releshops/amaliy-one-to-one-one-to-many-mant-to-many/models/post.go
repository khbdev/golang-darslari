package models

import "gorm.io/gorm"


 type Posts struct{
	gorm.Model
	UserID uint
	Content string
	Tags []Tag `gorm:"many2many:post_tags;"`
 }