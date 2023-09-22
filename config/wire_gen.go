// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package config

import (
	"GoGin-API-Base/api/handlers"
	"GoGin-API-Base/repository"
	"GoGin-API-Base/services"
	"github.com/google/wire"
)

// Injectors from injector.go:

func Init() *Initialization {
	gormDB := ConnectToDB()
	userRepositoryImpl := repository.UserRepositoryInit(gormDB)
	userServiceImpl := services.UserServiceInit(userRepositoryImpl)
	userHandlerImpl := handlers.UserHandlerInit(userServiceImpl)
	initialization := NewInitialization(userRepositoryImpl, userServiceImpl, userHandlerImpl)
	return initialization
}

// injector.go:

var db = wire.NewSet(ConnectToDB)

var userServiceSet = wire.NewSet(services.UserServiceInit, wire.Bind(new(services.UserService), new(*services.UserServiceImpl)))

var userRepoSet = wire.NewSet(repository.UserRepositoryInit, wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)))

var userHdlerSet = wire.NewSet(handlers.UserHandlerInit, wire.Bind(new(handlers.UserHandler), new(*handlers.UserHandlerImpl)))
