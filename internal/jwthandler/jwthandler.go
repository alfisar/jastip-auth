package jwthandler

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtHandler struct {
	Secret string
}

func GetJwt() *JwtHandler {
	return &JwtHandler{}
}

func (obj *JwtHandler) GetToken(lifepan time.Duration, id int) (token string, err error) {
	exp := time.Now().Add(lifepan)
	claims := jwt.MapClaims{
		"exp":     exp.Unix(),
		"iat":     time.Now().Unix(),
		"nbf":     time.Now().Unix(),
		"user_id": id,
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = jwtToken.SignedString([]byte(obj.Secret))
	return
}

func (obj *JwtHandler) ValidationToken(token string) (isValid bool, claim jwt.MapClaims, err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("")
			fmt.Printf("%s", r)
			err = fmt.Errorf("token is invalid")
			return
		}
	}()
	jwttoken, errData := jwt.Parse(token, func(jwttoken *jwt.Token) (interface{}, error) {
		if _, ok := jwttoken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		obj.Secret = os.Getenv("JWT_SECRET")
		return []byte(obj.Secret), nil
	})

	if errData != nil {
		if ve, ok := errData.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				err = fmt.Errorf("Token is expired")
				return
			}
		}
		err = fmt.Errorf("Token is Invalid")
		return
	}

	claim, ok := jwttoken.Claims.(jwt.MapClaims)
	if !ok {
		err = fmt.Errorf("cannot assert jwt payload to MapClaims")
		return
	}

	if !jwttoken.Valid {
		err = fmt.Errorf("Token is Invalid")
		return
	}

	isValid = true
	return

}
