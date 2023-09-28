package models

import (
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Title    string `form:"title"`
	Caption  string `form:"caption"`
	PhotoUrl string `form:"photo_url"`
	UserID   uint   
}
