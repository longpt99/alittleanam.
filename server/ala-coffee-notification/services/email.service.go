package services

import (
	"ala-coffee-notification/configs"
	"os"
	"text/template"

	"bytes"
	"fmt"
	"net/smtp"
)

type EmailService struct {
	auth smtp.Auth
	addr string
}

func InitEmailService() *EmailService {
	// Sender data.
	username := configs.Env.Email.Username
	password := configs.Env.Email.Password

	// smtp server configuration.
	smtpHost := configs.Env.Email.Host
	smtpPort := configs.Env.Email.Port

	auth := smtp.PlainAuth("", username, password, smtpHost)

	return &EmailService{
		auth: auth,
		addr: fmt.Sprintf(`%s:%s`, smtpHost, smtpPort),
	}
}

func (s *EmailService) SendEmail(from, to, name string) {
	pwd, _ := os.Getwd()

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Verify new account \n%s\n\n", mimeHeaders)))

	t, _ := template.ParseFiles(fmt.Sprintf("%s/assets/templates/verify-account.html", pwd))

	_ = t.Execute(&body, struct {
		Name string
		Link string
	}{
		Name: name,
		Link: "https://google.com",
	})

	err := smtp.SendMail(s.addr, s.auth, from, []string{to}, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Email Sent!")
}
