//go:build wireinject

package wire

import (
	"github.com/VanThen60hz/GoShop/internal/controller"
	"github.com/VanThen60hz/GoShop/internal/repo"
	"github.com/VanThen60hz/GoShop/internal/service"
	"github.com/google/wire"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepo,
		service.NewUserService,
		controller.NewUserController,
	)

	return new(controller.UserController), nil
}
