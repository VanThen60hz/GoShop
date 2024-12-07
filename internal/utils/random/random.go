package random

import (
	"math/rand"
	"time"
)

func GenerateSixDigitsOtp() int {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	otp := rng.Intn(900000) + 100000
	return otp
}
