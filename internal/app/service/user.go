package service

import (
	"api-service-template/internal/app/adapter"
	"api-service-template/internal/infra"
	"context"

	"gorm.io/gorm"
)

// UserService 用户服务
type UserService struct {
	User adapter.UserRepository
}

// NewUserService 构造函数
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		User: infra.NewUserDBRepository(db),
	}
}

// CheckUser 检查用户
func (s *UserService) CheckUser(ctx context.Context) error {
	return nil
}
