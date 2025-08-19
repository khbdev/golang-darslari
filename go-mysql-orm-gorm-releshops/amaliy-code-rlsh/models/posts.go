package models

import "gorm.io/gorm"


type Post struct {
	gorm.Model
	Content string  `gorm:"type:text;not null" validate:"required"`
	UserID  *uint   `gorm:"index"` // ✅ SET NULL ishlashi uchun *uint
	User    User    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	Tags    []Tags  `gorm:"many2many:post_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}