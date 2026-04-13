package app

import (
	"api-service-template/internal/app/adapter"
	"api-service-template/internal/app/command"

	"gorm.io/gorm"
)

type (
	// Application 业务聚合
	Application struct {
		Commands     *Commands
		Repositories *Repositories
	}

	// Commands 命令聚合
	Commands struct {
		CreateUser          *command.CreateUserHandler
		UpdateUser          *command.UpdateUserHandler
	}

	// Repositories 仓库聚合
	Repositories struct {
		User adapter.UserRepository
	}
)

// NewApplication 构造函数
func NewApplication(db *gorm.DB) *Application {
	return &Application{
		Commands:     InitCommands(db),
		Repositories: InitDBRepository(db),
	}
}
