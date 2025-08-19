package models

import "gorm.io/gorm"



type User struct {
	gorm.Model
	Name string
	Profile Profile `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}