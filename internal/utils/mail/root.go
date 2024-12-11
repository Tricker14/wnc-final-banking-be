package mail

import (
	"crypto/rand"
	"io"
	"strconv"

	"github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal/utils/env"
	"gopkg.in/gomail.v2"
)

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func GenerateOTP(max int) string {
    b := make([]byte, max)
    n, err := io.ReadAtLeast(rand.Reader, b, max)
    if n != max {
        panic(err)
    }
    for i := 0; i < len(b); i++ {
        b[i] = table[int(b[i])%len(table)]
    }
    return string(b)
}

func SendEmail(to, subject, body string) error {
    
    mailHost, _ := env.GetEnv("MAIL_HOST")
    mailPortStr, _ := env.GetEnv("MAIL_PORT")
    mailUsername, _ := env.GetEnv("MAIL_USERNAME")
    mailPassword, _ := env.GetEnv("MAIL_PASSWORD")
    mailFrom, _ := env.GetEnv("MAIL_FROM")

    mailPort, _ := strconv.Atoi(mailPortStr)

    m := gomail.NewMessage()
    m.SetHeader("From", mailFrom)
    m.SetHeader("To", to)
    m.SetHeader("Subject", subject)
    m.SetBody("text/html", body)

    d := gomail.NewDialer(mailHost, mailPort, mailUsername, mailPassword)
    return d.DialAndSend(m)
}