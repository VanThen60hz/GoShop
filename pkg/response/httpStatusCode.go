package response

const (
	ErrCodeSuccess      = 20001
	ErrCodeParamInvalid = 20003
	ErrInvalidToken     = 30001
	ErrInvalidOTP       = 30002
	ErrSendEmailOTP     = 30003

	// Register Code
	ErrCodeUserHasExists = 50001 // User has already registered
)

var msg = map[int]string{
	ErrCodeSuccess:      "Success",
	ErrCodeParamInvalid: "Email is invalid",
	ErrInvalidToken:     "Invalid Token",
	ErrInvalidOTP:       "Invalid OTP",
	ErrSendEmailOTP:     "Failed to send email OTP",

	// Register Code
	ErrCodeUserHasExists: "User has already registered",
}
