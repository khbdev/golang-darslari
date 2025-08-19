package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name    string    `gorm:"not null" validate:"required"`
	Profile *Profile  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // <- pointer bo'lishi kerak!
	Posts   []Post    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}