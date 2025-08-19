package models

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	Bio    string `gorm:"type:text" validate:"max=1000"`
	UserID *uint  `gorm:"unique"` // ✅ *uint bo'lishi kerak, NOT NULL emas!

	User   User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}