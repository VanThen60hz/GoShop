package repo

import (
	"fmt"
	"time"

	"github.com/VanThen60hz/GoShop/global"
)

type IUserAuthRepo interface {
	AddOTP(email string, otp int, expirationTime int64) error
}

type userAuthRepo struct{}

func NewUserAuthRepo() IUserAuthRepo {
	return &userAuthRepo{}
}

// AddOTP implements IUserAuthRepo.
func (u *userAuthRepo) AddOTP(email string, otp int, expirationTime int64) error {
	key := fmt.Sprintf("usr:%s:otp", email)
	return global.Rdb.Set(ctx, key, otp, time.Duration(expirationTime)).Err()
}
