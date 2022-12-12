package services

import (
	"auth_microservice/config"
	"fmt"
	mail "github.com/xhit/go-simple-mail/v2"
	"math/rand"
	"os"
	"time"
)

func Send(code int, username string) bool {
	address := mail.NewMSG()
	address.SetFrom("Survivor Coders <narjeswehbi04@gmail.com>")
	address.AddTo(username)
	address.SetSubject("Verify your email")
	b, e := os.ReadFile("msg.html")
	if e != nil {
		fmt.Println("failed to load html file !! ")
		return false
	}

	address.SetBody(mail.TextHTML, fmt.Sprintf(string(b), code))

	err := address.Send(config.SmtpClient)
	if err != nil {
		fmt.Println("error mail.go")
		return false
	}
	return true
}

func GenerateCode() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(9999-1000) + 1000
}
