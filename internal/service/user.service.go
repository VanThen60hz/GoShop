package service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/VanThen60hz/GoShop/internal/repo"
	"github.com/VanThen60hz/GoShop/internal/utils/cryto"
	"github.com/VanThen60hz/GoShop/internal/utils/random"
	sendto "github.com/VanThen60hz/GoShop/internal/utils/sendTo"
	"github.com/VanThen60hz/GoShop/pkg/response"
)

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo     repo.IUserRepo
	userAuthRepo repo.IUserAuthRepo
}

func NewUserService(userRepo repo.IUserRepo, userAuthRepo repo.IUserAuthRepo) IUserService {
	return &userService{
		userRepo:     userRepo,
		userAuthRepo: userAuthRepo,
	}
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {
	hashEmail := cryto.GetHash(email)
	fmt.Println("Hash Email: ", hashEmail)

	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExists
	}

	otp := random.GenerateSixDigitsOtp()
	if purpose == "TEST_USER" {
		otp = 123456
	}

	fmt.Println("OTP: ", otp)

	err := us.userAuthRepo.AddOTP(hashEmail, otp, int64(10*time.Minute))
	if err != nil {
		return response.ErrInvalidOTP
	}

	err = sendto.SendTemplateEmailOtp([]string{email}, "nguyenvthang2409@gmail.com", "otp-auth.html", map[string]interface{}{
		"otp": strconv.Itoa(otp),
	})
	if err != nil {
		return response.ErrSendEmailOTP
	}

	return response.ErrCodeSuccess
}
