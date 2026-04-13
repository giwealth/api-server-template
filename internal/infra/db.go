package infra

import (
	"gorm.io/gorm"
)

// Init 初始化
func Init(db *gorm.DB) error {
	return db.AutoMigrate(
		&userRow{},
	)
}
