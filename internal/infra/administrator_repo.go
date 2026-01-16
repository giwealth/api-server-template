package infra

import (
	"api-service-template/internal/domain"
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type administratorRow struct {
	ID       int64
	Username string
	Password string
}

type administratorRows []administratorRow

// AdministratorDBRepository 存储
type AdministratorDBRepository struct {
	db *gorm.DB
}

// NewAdministratorDBRepository 构造函数
func NewAdministratorDBRepository(db *gorm.DB) *AdministratorDBRepository {
	return &AdministratorDBRepository{db}
}

// TableName 表名称
func (administratorRow) TableName() string {
	return "administrators"
}

// Create 创建
func (r *AdministratorDBRepository) Create(ctx context.Context, administrator *domain.Administrator) error {
	row := administratorRow{
		ID:       administrator.ID,
		Username: administrator.Username,
		Password: administrator.Password,
	}

	if err := r.db.WithContext(ctx).Create(&row).Error; err != nil {
		return errors.WithStack(err)
	}

	administrator.ID = row.ID

	return nil
}

// Delete 删除
func (r *AdministratorDBRepository) Delete(ctx context.Context, id int64) error {
	err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&administratorRow{}).Error
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Update 更改
func (r *AdministratorDBRepository) Update(ctx context.Context, administrator *domain.Administrator) error {
	row := administratorRow{
		ID:       administrator.ID,
		Username: administrator.Username,
		Password: administrator.Password,
	}

	if err := r.db.WithContext(ctx).Model(&administratorRow{}).Where("id = ?", administrator.ID).Updates(&row).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Find 查找
func (r *AdministratorDBRepository) Find(ctx context.Context, id int64) (*domain.Administrator, error) {
	var row administratorRow
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&row).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrAdministratorNotFound
		}
		return nil, errors.WithStack(err)
	}

	return row.toDomainAdministrator(), nil
}

// List GetForPage 分页获取
func (r *AdministratorDBRepository) List(ctx context.Context, page, limit int) ([]*domain.Administrator, int64, error) {
	var rows administratorRows
	var total int64
	offset := limit * (page - 1)
	if err := r.db.WithContext(ctx).Order("id desc").Limit(limit).Offset(offset).Find(&rows).Error; err != nil {
		return nil, total, errors.WithStack(err)
	}

	if err := r.db.WithContext(ctx).Model(&administratorRow{}).Count(&total).Error; err != nil {
		return nil, total, errors.WithStack(err)
	}

	return rows.toDomainAdministrators(), total, nil
}

func (row administratorRow) toDomainAdministrator() *domain.Administrator {
	return &domain.Administrator{
		ID:       row.ID,
		Username: row.Username,
		Password: row.Password,
	}
}

func (rows administratorRows) toDomainAdministrators() []*domain.Administrator {
	l := make([]*domain.Administrator, 0, len(rows))

	for _, row := range rows {
		v := row.toDomainAdministrator()
		l = append(l, v)
	}

	return l
}
