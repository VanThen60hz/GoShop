package repo

import (
	"github.com/VanThen60hz/GoShop/global"
	"github.com/VanThen60hz/GoShop/internal/model"
)

// type UserRepo struct{}

// func NewUserRepo() *UserRepo {
// 	return &UserRepo{}
// }

// func (ur *UserRepo) GetInfoUser() string {
// 	return "Golang"
// }

// INTERFACE VERSION
type IUserRepo interface {
	GetUserByEmail(email string) bool
}

type userRepo struct{}

func NewUserRepo() IUserRepo {
	return &userRepo{}
}

// GetUserByEmail implements IUserRepo.
func (u *userRepo) GetUserByEmail(email string) bool {
	// SELECT * FROM user WHERE email = email = '??' ORDER BY email
	row := global.Mdb.Table(TableNameGoCrmUser).Where("user_email = ?", email).First(model.GoCrmUser{}).RowsAffected
	return row != NumberNull
}
