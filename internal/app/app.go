package app

import (
	"api-service-template/internal/app/adapter"
	"api-service-template/internal/app/command"
	"api-service-template/internal/option"
)

type (
	// Application 业务聚合
	Application struct {
		Commands     *Commands
		Repositories *Repositories
		Opt          *option.Options
	}

	// Commands 命令聚合
	Commands struct {
		CreateUser *command.CreateUserHandler
		UpdateUser *command.UpdateUserHandler
	}

	// Repositories 仓库聚合
	Repositories struct {
		User adapter.UserRepository
	}
)

// NewApplication 构造函数
func NewApplication(opt *option.Options) *Application {
	db := opt.GetDB()
	return &Application{
		Commands:     InitCommands(db),
		Repositories: InitDBRepository(db),
		Opt:          opt,
	}
}
