package repo

import (
	"github.com/VanThen60hz/GoShop/global"
	"github.com/VanThen60hz/GoShop/internal/database"
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

type userRepo struct {
	sqlc *database.Queries
}

func NewUserRepo() IUserRepo {
	return &userRepo{
		sqlc: database.New(global.Mdbc),
	}
}

// GetUserByEmail implements IUserRepo.
func (ur *userRepo) GetUserByEmail(email string) bool {
	// SELECT * FROM user WHERE email = email = '??' ORDER BY email
	// row := global.Mdb.Table(TableNameGoCrmUser).Where("user_email = ?", email).First(model.GoCrmUser{}).RowsAffected
	user, err := ur.sqlc.GetUserByEmailSQLC(ctx, email)
	if err != nil {
		return false
	}

	return user.UsrID != 0
}
