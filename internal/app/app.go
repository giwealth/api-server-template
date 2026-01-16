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
		DeleteUser          *command.DeleteUserHandler
		UpdateUser          *command.UpdateUserHandler
		FindUser            *command.FindUserHandler
		ListUser            *command.ListUserHandler
		CreateAdministrator *command.CreateAdministratorHandler
		DeleteAdministrator *command.DeleteAdministratorHandler
		UpdateAdministrator *command.UpdateAdministratorHandler
		FindAdministrator   *command.FindAdministratorHandler
		ListAdministrator   *command.ListAdministratorHandler
	}

	// Repositories 仓库聚合
	Repositories struct {
		User          adapter.UserRepository
		Administrator adapter.AdministratorRepository
	}
)

// NewApplication 构造函数
func NewApplication(db *gorm.DB) *Application {
	return &Application{
		Commands:     InitCommands(db),
		Repositories: InitDBRepository(db),
	}
}
