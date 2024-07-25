package handler

import (
	"github.com/soulwill1/estudos/configs"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitializeHandler() {
	db = configs.GetMySQL()
}