package command

import (
	"api-service-template/internal/app/adapter"
	"api-service-template/internal/app/service"
	"api-service-template/internal/domain"
	"context"
)

type (
	// CreateUserHandler 示例
	CreateUserHandler struct {
		User      adapter.UserRepository
		CheckUser service.UserService
	}
	// DeleteUserHandler 示例
	DeleteUserHandler struct {
		User adapter.UserRepository
	}
	// UpdateUserHandler 示例
	UpdateUserHandler struct {
		User adapter.UserRepository
	}
	// FindUserHandler 示例
	FindUserHandler struct {
		User adapter.UserRepository
	}
	// ListUserHandler 示例
	ListUserHandler struct {
		User adapter.UserRepository
	}
)

// Handle 执行
func (h CreateUserHandler) Handle(ctx context.Context, user *domain.User) error {
	// TODO
	if err := h.CheckUser.CheckUser(ctx); err != nil {
		return err
	}
	return h.User.Create(ctx, user)
}

// Handle 执行
func (h DeleteUserHandler) Handle(ctx context.Context, id int64) error {
	// TODO
	return h.User.Delete(ctx, id)
}

// Handle 执行
func (h UpdateUserHandler) Handle(ctx context.Context, user *domain.User) error {
	// TODO
	return h.User.Update(ctx, user)
}

// Handle 执行
func (h FindUserHandler) Handle(ctx context.Context, id int64) (*domain.User, error) {
	// TODO
	return h.User.Find(ctx, id)
}

// Handle 执行查的，返回列表，总数
func (h ListUserHandler) Handle(ctx context.Context, page, limit int) ([]*domain.User, int64, error) {
	// TODO
	return h.User.List(ctx, page, limit)
}
