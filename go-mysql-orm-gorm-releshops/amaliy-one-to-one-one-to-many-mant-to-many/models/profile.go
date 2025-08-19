package models

import "gorm.io/gorm"


type Profile struct{
	gorm.Model
	UserID uint
	Bio string
}