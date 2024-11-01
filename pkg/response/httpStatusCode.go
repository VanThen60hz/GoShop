package response

const (
	ErrCodeSuccess      = 20001
	ErrCodeParamInvalid = 20003
	ErrInvalidToken     = 30001

	// Register Code
	ErrCodeUserHasExists = 50001 // User has already registered
)

var msg = map[int]string{
	ErrCodeSuccess:      "Success",
	ErrCodeParamInvalid: "Email is invalid",
	ErrInvalidToken:     "Invalid token",

	// Register Code
	ErrCodeUserHasExists: "User has already registered",
}
