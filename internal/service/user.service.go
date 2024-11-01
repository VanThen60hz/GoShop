package service

import (
	"github.com/VanThen60hz/GoShop/internal/repo"
	"github.com/VanThen60hz/GoShop/pkg/response"
)

// import "github.com/VanThen60hz/GoShop/internal/repo"

// type UserService struct {
// 	UserRepo *repo.UserRepo
// }

// func NewUserService() *UserService {
// 	return &UserService{
// 		UserRepo: repo.NewUserRepo(),
// 	}
// }

// func (us *UserService) GetInfoUser() string {
// 	return us.UserRepo.GetInfoUser()
// }

// INTERFACE VERSION
type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo repo.IUserRepo
}

func NewUserService(userRepo repo.IUserRepo) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExists
	}

	return response.ErrCodeSuccess
}
