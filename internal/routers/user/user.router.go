package user

import (
	"github.com/VanThen60hz/GoShop/internal/wire"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	// userRepo := repo.NewUserRepo()
	// userService := service.NewUserService(userRepo)
	// userHandlerNonDependency := controller.NewUserController(userService)

	userController, _ := wire.InitUserRouterHandler()

	// public router
	userRouterPublic := router.Group("/user")
	{
		userRouterPublic.POST("/register", userController.Register)
		// userRouterPublic.POST("/otp")
	}

	// private router
	userRouterPrivate := router.Group("/user")
	// userRouterPrivate.Use(limiter())
	// userRouterPrivate.Use(Authen())
	// userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.GET("/get-info")
	}
}
