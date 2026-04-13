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
	// UpdateUserHandler 示例
	UpdateUserHandler struct {
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
func (h UpdateUserHandler) Handle(ctx context.Context, user *domain.User) error {
	// TODO
	return h.User.Update(ctx, user)
}
