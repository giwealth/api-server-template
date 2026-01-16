package infra

import (
	"api-service-template/internal/domain"
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type userRow struct {
	ID      int64
	Name    string
	Age     int
	Address string
}

type userRows []userRow

// UserDBRepository 存储
type UserDBRepository struct {
	db *gorm.DB
}

// NewUserDBRepository 构造函数
func NewUserDBRepository(db *gorm.DB) *UserDBRepository {
	return &UserDBRepository{db}
}

// TableName 表名称
func (userRow) TableName() string {
	return "users"
}

// Create 创建
func (r *UserDBRepository) Create(ctx context.Context, user *domain.User) error {
	row := userRow{
		ID:      user.ID,
		Name:    user.Name,
		Age:     user.Age,
		Address: user.Address,
	}

	if err := r.db.WithContext(ctx).Create(&row).Error; err != nil {
		return errors.WithStack(err)
	}

	user.ID = row.ID

	return nil
}

// Delete 删除
func (r *UserDBRepository) Delete(ctx context.Context, id int64) error {
	err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&userRow{}).Error
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Update 更改
func (r *UserDBRepository) Update(ctx context.Context, user *domain.User) error {
	row := userRow{
		ID:      user.ID,
		Name:    user.Name,
		Age:     user.Age,
		Address: user.Address,
	}

	if err := r.db.WithContext(ctx).Model(&userRow{}).Where("id = ?", user.ID).Updates(&row).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Find 查找
func (r *UserDBRepository) Find(ctx context.Context, id int64) (*domain.User, error) {
	var row userRow
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&row).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrUserNotFound
		}
		return nil, errors.WithStack(err)
	}

	return row.toDomainUser(), nil
}

// List GetForPage 分页获取
func (r *UserDBRepository) List(ctx context.Context, page, limit int) ([]*domain.User, int64, error) {
	var rows userRows
	var total int64
	offset := limit * (page - 1)
	if err := r.db.WithContext(ctx).Order("id desc").Limit(limit).Offset(offset).Find(&rows).Error; err != nil {
		return nil, total, errors.WithStack(err)
	}

	if err := r.db.WithContext(ctx).Model(&userRow{}).Count(&total).Error; err != nil {
		return nil, total, errors.WithStack(err)
	}

	return rows.toDomainUsers(), total, nil
}

func (row userRow) toDomainUser() *domain.User {
	return &domain.User{
		ID:      row.ID,
		Name:    row.Name,
		Age:     row.Age,
		Address: row.Address,
	}
}

func (rows userRows) toDomainUsers() []*domain.User {
	l := make([]*domain.User, 0, len(rows))

	for _, row := range rows {
		v := row.toDomainUser()
		l = append(l, v)
	}

	return l
}
