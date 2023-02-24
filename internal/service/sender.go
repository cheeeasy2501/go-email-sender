package service

import (
	"fmt"
	"net/smtp"
)

type ISenderService interface {
	SendMail() (bool, error)
}

type SenderService struct {
}

func NewSenderService() *SenderService {
	return &SenderService{}
}

func (s *SenderService) SendMail() (bool, error) {
	/**TODO: example send mail*/
	from := "john.doe@example.com"

	user := "9c1d45eaf7af5b"
	password := "ad62926fa75d0f"

	to := []string{
		"roger.roe@example.com",
	}

	addr := "smtp.mailtrap.io:2525"
	host := "smtp.mailtrap.io"

	msg := []byte("From: john.doe@example.com\r\n" +
		"To: roger.roe@example.com\r\n" +
		"Subject: Test mail\r\n\r\n" +
		"Email body\r\n")

	auth := smtp.PlainAuth("", user, password, host)

	err := smtp.SendMail(addr, auth, from, to, msg)

	if err != nil {
		return false, err
	}

	fmt.Println("Email sent successfully")
	return true, nil
}
