package database

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Option 数据库选项
type Option struct {
	Driver       string `yaml:"driver"`
	DSN          string `yaml:"dsn"`
	MaxOpenConns int    `yaml:"maxOpenConns"`
	MaxIdleConns int    `yaml:"maxIdleConns"`
}

// NewDB 创建数据库连接
func NewDB(opt Option) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(opt.DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if opt.MaxOpenConns > 0 {
		sqlDB.SetMaxOpenConns(opt.MaxOpenConns)
	}
	if opt.MaxIdleConns > 0 {
		sqlDB.SetMaxIdleConns(opt.MaxIdleConns)
	}

	if os.Getenv("LEVEL") == "debug" {
		db = db.Debug()
	}

	return db, nil
}
