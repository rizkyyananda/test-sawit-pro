package config

import (
	"gorm.io/gorm"
	"test_sawit_pro/entity"
)

func Migration(db *gorm.DB) {
	err := db.AutoMigrate(&entity.Tree{}, &entity.Estate{})
	if err != nil {
		return
	}
}
