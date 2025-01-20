package general

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func GetRandomOTP(length int) (OTP string, err error) {
	max := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(length)), nil)

	num, err := rand.Int(rand.Reader, max)
	if err != nil {
		fmt.Println("get random OTP is Error : " + err.Error())
		return
	}

	OTP = num.String()
	if len(OTP) < length {
		OTP = "0" + OTP
	}

	OTP = OTP[:length]
	return
}
