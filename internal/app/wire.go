//go:build wireinject
// +build wireinject

package app

import (
	"api-service-template/internal/app/adapter"
	"api-service-template/internal/app/command"
	"api-service-template/internal/app/service"
	"api-service-template/internal/infra"

	"github.com/google/wire"
	"gorm.io/gorm"
)

var (
	dbRepositorySet = wire.NewSet(
		wire.NewSet(
			infra.NewUserDBRepository,
			wire.Bind(new(adapter.UserRepository), new(*infra.UserDBRepository)),
			infra.NewAdministratorDBRepository,
			wire.Bind(new(adapter.AdministratorRepository), new(*infra.AdministratorDBRepository)),
		),
	)

	serviceSet = wire.NewSet(
		wire.Struct(new(service.UserService), "*"),
	)

	dbRepositoryProvider = wire.NewSet(
		dbRepositorySet,
		wire.Struct(new(Repositories), "*"),
	)

	commandsProvider = wire.NewSet(
		serviceSet,
		dbRepositorySet,
		wire.Struct(new(command.CreateUserHandler), "*"),
		wire.Struct(new(command.DeleteUserHandler), "*"),
		wire.Struct(new(command.UpdateUserHandler), "*"),
		wire.Struct(new(command.FindUserHandler), "*"),
		wire.Struct(new(command.ListUserHandler), "*"),
		wire.Struct(new(command.CreateAdministratorHandler), "*"),
		wire.Struct(new(command.DeleteAdministratorHandler), "*"),
		wire.Struct(new(command.UpdateAdministratorHandler), "*"),
		wire.Struct(new(command.FindAdministratorHandler), "*"),
		wire.Struct(new(command.ListAdministratorHandler), "*"),
		wire.Struct(new(Commands), "*"),
	)
)

func InitDBRepository(db *gorm.DB) *Repositories {
	wire.Build(dbRepositoryProvider)
	return &Repositories{}
}

func InitCommands(db *gorm.DB) *Commands {
	wire.Build(commandsProvider)
	return &Commands{}
}
