package repo

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
	return true
}
