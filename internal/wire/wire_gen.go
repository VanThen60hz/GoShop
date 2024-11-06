// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/VanThen60hz/GoShop/internal/controller"
	"github.com/VanThen60hz/GoShop/internal/repo"
	"github.com/VanThen60hz/GoShop/internal/service"
)

// Injectors from user.wire.go:

func InitUserRouterHandler() (*controller.UserController, error) {
	iUserRepo := repo.NewUserRepo()
	iUserService := service.NewUserService(iUserRepo)
	userController := controller.NewUserController(iUserService)
	return userController, nil
}