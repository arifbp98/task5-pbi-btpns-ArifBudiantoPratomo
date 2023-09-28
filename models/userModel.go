package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string  `form:"username" valid:"required"`
	Email    string  `form:"email" gorm:"unique" valid:"required,email"`
	Password string  `form:"password" valid:"required"`
	Photos   []Photo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
