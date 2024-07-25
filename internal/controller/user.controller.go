package controller

import (
	"github.com/VanThen60hz/GoShop/internal/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	// if err != nil {
	// 	return response.ErrResponse(c, 20003, "No need!")
	// }

	// return response.SuccessResponse(c, 20001, []string{"cr7", "ronaldo", "m10"})
}
