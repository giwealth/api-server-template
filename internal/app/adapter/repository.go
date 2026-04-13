package adapter

import (
	"api-service-template/internal/domain"
	"context"
)

type (
	// UserRepository 用户存储
	UserRepository interface {
		Create(ctx context.Context, user *domain.User) error
		Delete(ctx context.Context, id int64) error
		Update(ctx context.Context, user *domain.User) error
		Find(ctx context.Context, id int64) (*domain.User, error)
		List(ctx context.Context, page, limit int) ([]*domain.User, int64, error)
	}
)
