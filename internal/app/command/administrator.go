package command

import (
	"api-service-template/internal/app/adapter"
	"api-service-template/internal/domain"
	"context"
)

type (
	// CreateAdministratorHandler 示例
	CreateAdministratorHandler struct {
		Administrator adapter.AdministratorRepository
	}
	// DeleteAdministratorHandler 示例
	DeleteAdministratorHandler struct {
		Administrator adapter.AdministratorRepository
	}
	// UpdateAdministratorHandler 示例
	UpdateAdministratorHandler struct {
		Administrator adapter.AdministratorRepository
	}
	// FindAdministratorHandler 示例
	FindAdministratorHandler struct {
		Administrator adapter.AdministratorRepository
	}
	// ListAdministratorHandler 示例
	ListAdministratorHandler struct {
		Administrator adapter.AdministratorRepository
	}
)

// Handle 执行
func (h CreateAdministratorHandler) Handle(ctx context.Context, administrator *domain.Administrator) error {
	// TODO
	return h.Administrator.Create(ctx, administrator)
}

// Handle 执行
func (h DeleteAdministratorHandler) Handle(ctx context.Context, id int64) error {
	// TODO
	return h.Administrator.Delete(ctx, id)
}

// Handle 执行
func (h UpdateAdministratorHandler) Handle(ctx context.Context, administrator *domain.Administrator) error {
	// TODO
	return h.Administrator.Update(ctx, administrator)
}

// Handle 执行
func (h FindAdministratorHandler) Handle(ctx context.Context, id int64) (*domain.Administrator, error) {
	// TODO
	return h.Administrator.Find(ctx, id)
}

// Handle 执行查的，返回列表，总数
func (h ListAdministratorHandler) Handle(ctx context.Context, page, limit int) ([]*domain.Administrator, int64, error) {
	// TODO
	return h.Administrator.List(ctx, page, limit)
}
