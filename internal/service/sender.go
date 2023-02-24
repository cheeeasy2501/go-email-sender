package service

import (
	"fmt"
	"net/smtp"

	"github.com/cheeeasy2501/go-email-sender/config"
	"github.com/cheeeasy2501/go-email-sender/internal/dto"
)

type ISenderService interface {
	SendMail(d dto.EmailDTO) (bool, error)
}

type SenderService struct {
	cfg *config.Config
}

func NewSenderService(cfg *config.Config) *SenderService {
	return &SenderService{
		cfg: cfg,
	}
}

func (s *SenderService) SendMail(d dto.EmailDTO) (bool, error) {
	addr := s.cfg.Mail().GetAddress()
	host := s.cfg.Mail().GetHost()
	from := s.cfg.Mail().GetAddressFrom()
	user := s.cfg.Mail().GetUsername()
	password := s.cfg.Mail().GetPassword()

	// TODO: придумать как собрать msg по RFC. Возможно builder
	msg := []byte("From: " + from + "\r\n" +
		"To: " + d.To()[0] + "\r\n" +
		"Subject:" + d.Subject() + "\r\n\r\n" +
		d.Body() + "\r\n")

	auth := smtp.PlainAuth("", user, password, host)

	err := smtp.SendMail(addr, auth, from, d.To(), msg)
	if err != nil {
		return false, err
	}
	// TODO: заменить на нормальный logger
	fmt.Println("Email sent successfully")
	return true, nil
}
