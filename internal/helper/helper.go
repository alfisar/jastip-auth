package helper

import (
	"fmt"
	"jastip/config"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendEmailOTP(pooldata *config.Config, email string, fullname string, OTP string) (err error) {
	body := "Hello " + fullname + " , thank u for the registration, so this the OTP for confirmation your registration, OTP : " + OTP

	_port, _ := strconv.Atoi(pooldata.SMTP.Port)

	mailer := pooldata.SMTP.Mailer
	mailer.SetHeader("From", pooldata.SMTP.From)
	mailer.SetHeader("To", email)
	mailer.SetHeader("Subject", "Registration Jastip.in")
	mailer.SetBody("text/html", body)

	dialer := gomail.NewDialer(
		pooldata.SMTP.Host,
		_port,
		pooldata.SMTP.User,
		pooldata.SMTP.Pass,
	)

	err = dialer.DialAndSend(mailer)

	if err != nil {
		fmt.Println("Sending email is Error : " + err.Error())
		return
	}

	fmt.Println("Sending email is success")
	return
}
