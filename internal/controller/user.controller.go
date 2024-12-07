package controller

import (
	"fmt"

	"github.com/VanThen60hz/GoShop/internal/service"
	"github.com/VanThen60hz/GoShop/internal/vo"
	"github.com/VanThen60hz/GoShop/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(
	userService service.IUserService,
) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	var params vo.UserRegistrationRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.ErrResponse(c, response.ErrCodeParamInvalid, err.Error())
		return
	}
	fmt.Println("Email Params: ", params.Email)
	res := uc.userService.Register(params.Email, params.Purpose)
	response.SuccessResponse(c, res, nil)
}

// func (uc *UserController) GetUserByID(c *gin.Context) {
// 	// if err != nil {
// 	// 	return response.ErrResponse(c, 20003, "No need!")
// 	// }

// 	// return response.SuccessResponse(c, 20001, []string{"cr7", "ronaldo", "m10"})
// }
