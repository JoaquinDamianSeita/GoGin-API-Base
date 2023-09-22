// go:build wireinject
//go:build wireinject
// +build wireinject

package config

import (
	"GoGin-API-Base/api/handlers"
	"GoGin-API-Base/repository"
	"GoGin-API-Base/services"

	"github.com/google/wire"
)

var db = wire.NewSet(ConnectToDB)

var userServiceSet = wire.NewSet(services.UserServiceInit,
	wire.Bind(new(services.UserService), new(*services.UserServiceImpl)),
)

var userRepoSet = wire.NewSet(repository.UserRepositoryInit,
	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
)

var userHdlerSet = wire.NewSet(handlers.UserHandlerInit,
	wire.Bind(new(handlers.UserHandler), new(*handlers.UserHandlerImpl)),
)

func Init() *Initialization {
	wire.Build(NewInitialization, db, userHdlerSet, userServiceSet, userRepoSet)
	return nil
}
